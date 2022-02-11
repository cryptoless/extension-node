package service

import (
	"bytes"
	"context"
	"crypto/sha256"
	"extension-node/util/model"
	"fmt"
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
	if fnType.NumOut() > 3 {
		g.Log().Error("Registration Api:", fnType.String(), ",unsuport result>3")
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
func (a *callObj) CallAble(ctx context.Context, args []reflect.Value) *CallAble {
	ca := &CallAble{
		errPos: a.errPos,
	}
	// build fullargs
	ca.args = []reflect.Value{}
	if a.rcv.IsValid() {
		ca.args = append(ca.args, a.rcv)
	}
	if a.hasCtx {
		ca.args = append(ca.args, reflect.ValueOf(ctx))
	}
	// args
	ca.args = append(ca.args, args...)

	ca.fn = a.fn
	ca.obj = a

	return ca
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

	// bufa, err := json.Marshal(a)
	// fmt.Println(bufa)
	// if err != nil {
	// 	g.Log().Error(err)
	// 	return "", err
	// }

	// id := hash.Sum(buf.Bytes())
	// return string(id), nil
}

func (a *CallAble) Call(ctx context.Context) (res interface{}, errRes error) {

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
