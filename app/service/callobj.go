package service

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"extension-node/app/model"
	"fmt"
	"io"
	"reflect"
	"runtime"

	"github.com/gogf/gf/frame/g"
)

type callObj struct {
	fn  reflect.Value
	rcv reflect.Value

	argsType []reflect.Type

	isSub  bool
	hasCtx bool
	errPos int
}

func buildCallObj(fn reflect.Value, rcvVal reflect.Value, isSub bool) *callObj {
	call := &callObj{
		fn:  fn,
		rcv: rcvVal,
	}
	fnType := fn.Type()
	fistArg := 0
	if rcvVal.IsValid() {
		fistArg++
	}
	if fnType.NumIn() > fistArg && fnType.In(fistArg) == contextType {
		call.hasCtx = true
		fistArg++
	}
	// in
	for i := fistArg; i < fnType.NumIn(); i++ {
		argType := fnType.In(i)
		call.argsType = append(call.argsType, argType)
	}

	// out
	if fnType.NumOut() > 2 {
		g.Log().Error("Registration Api:", fnType.String(), ",unsupport result>2")
		return nil
	}
	if fnType.NumOut() == 0 {
		g.Log().Error("Registration Api:", fnType.String(), ",want error")
		return nil
	}
	if fnType.Out(0) == errType {
		call.errPos = 0
	} else {
		call.errPos = 1
	}
	call.isSub = isSub
	return call
}

func parseArray(decoder *json.Decoder, argsType []reflect.Type) ([]reflect.Value, error) {
	args := make([]reflect.Value, 0, len(argsType))
	for i := 0; decoder.More(); i++ {
		if i >= len(argsType) {
			return nil, fmt.Errorf("too many args, want %d args", len(argsType))
		}
		val := reflect.New(argsType[i])
		if err := decoder.Decode(val.Interface()); err != nil {
			g.Log().Error(err)
			return nil, err
		}

		if val.IsNil() && val.Kind() == reflect.Ptr {
			return nil, fmt.Errorf("missing value for args %d", i)
		}
		args = append(args, val.Elem())
	}
	_, err := decoder.Token()
	return args, err
}

func parseArgs(rawMsg json.RawMessage, argsType []reflect.Type) ([]reflect.Value, error) {
	decoder := json.NewDecoder(bytes.NewReader(rawMsg))
	var args []reflect.Value
	token, err := decoder.Token()

	switch {
	case err == io.EOF || token == nil && err == nil:
	case err != nil:
		return nil, err

	case token == json.Delim('['):
		if args, err = parseArray(decoder, argsType); err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("no array")
	}

	for i := len(args); i < len(argsType); i++ {
		if argsType[i].Kind() != reflect.Ptr && argsType[i].Kind() != reflect.Slice {
			return nil, fmt.Errorf("missing value for required argument %d", i)
		}
		args = append(args, reflect.Zero(argsType[i]))
	}
	return args, nil
}

func (a *callObj) CallAble(ctx context.Context, msg *model.JsonMessage) (ca *CallAble, errRes error) {
	//catch panic
	defer func() {
		if err := recover(); err != nil {
			const size = 64 << 10
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)]
			method := a.fn.Type().Name()
			g.Log().Error("methdo:" + method + " crashed: " + fmt.Sprintf("%v\n%s", err, buf))
			errRes = fmt.Errorf("methdo:"+method+" crashed:%+v", err)
		}
	}()
	//build args
	args, err := parseArgs(msg.Params, a.argsType)
	if err != nil {
		return nil, err
	}
	// build fullargs
	fullargs := []reflect.Value{}
	if a.rcv.IsValid() {
		fullargs = append(fullargs, a.rcv)
	}
	if a.hasCtx {
		fullargs = append(fullargs, reflect.ValueOf(ctx))
	}
	// fullargs
	fullargs = append(fullargs, args...)

	ca = &CallAble{
		ctx:    ctx,
		fn:     a.fn,
		args:   fullargs,
		errPos: a.errPos,
		msg:    msg,
		obj:    a,
	}
	// id
	hash := sha256.New()
	buf := bytes.Buffer{}
	buf.Write(msg.Id)
	buf.WriteString(msg.Method)
	ca.id = string(hash.Sum(buf.Bytes()))

	return ca, nil
}

type CallAble struct {
	ctx    context.Context
	fn     reflect.Value
	args   []reflect.Value
	errPos int
	msg    *model.JsonMessage
	id     string
	obj    *callObj
}

func (a *CallAble) Id() string {

	if a.id == "" {
		hash := sha256.New()
		buf := bytes.Buffer{}
		buf.Write(a.msg.Id)
		buf.WriteString(a.msg.Method)
		a.id = string(hash.Sum(buf.Bytes()))
	}
	return a.id

}

func (a *CallAble) Call() (res interface{}, errRes error) {

	//catch panic
	defer func() {
		if err := recover(); err != nil {
			const size = 64 << 10
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)]
			method := a.fn.Type().Name()
			g.Log().Error("methdo:" + method + " crashed: " + fmt.Sprintf("%v\n%s", err, buf))
			errRes = fmt.Errorf("methdo:"+method+" crashed:%+v", err)
		}
	}()

	//call
	rs := a.fn.Call(a.args)
	if len(rs) == 0 {
		return nil, nil
	}
	if a.errPos >= 0 && !rs[a.errPos].IsNil() {
		err := rs[a.errPos].Interface().(error)
		return reflect.Value{}, err
	}
	return rs[0].Interface(), nil
}
