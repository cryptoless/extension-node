package service

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"extension-node/app/service/eth"
	"extension-node/util/rpc"
	"fmt"
	"io"
	"path/filepath"
	"reflect"
	"runtime"
	"sync"
	"unicode"

	"github.com/gogf/gf/frame/g"
)

type callObj struct {
	fn  reflect.Value
	rcv reflect.Value

	argsType []reflect.Type
	hasCtx   bool
	errPos   int
}

func buildCallObj(fn reflect.Value, rcvVal reflect.Value) *callObj {
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
		g.Log().Error("Registration Api:", fnType.String(), ",unsuport multi-ret")
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
	return call
}

func (a *callObj) call(ctx context.Context, args []reflect.Value) (res interface{}, errRes error) {

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
	// build fullargs
	fullargs := []reflect.Value{}
	if a.rcv.IsValid() {
		fullargs = append(fullargs, a.rcv)
	}
	if a.hasCtx {
		fullargs = append(fullargs, reflect.ValueOf(ctx))
	}
	// args
	fullargs = append(fullargs, args...)
	//call
	rs := a.fn.Call(fullargs)
	if len(rs) == 0 {
		return nil, nil
	}
	if a.errPos >= 0 && !rs[a.errPos].IsNil() {
		err := rs[a.errPos].Interface().(error)
		return reflect.Value{}, err
	}
	g.Log().Debugf("call rs:%+v\n", rs)
	return rs[0].Interface(), nil
}

type serviceRegister struct {
	callList map[string]*callObj
	rcvList  map[string]interface{}
}

var Service *serviceRegister

var once sync.Once

func init() {
	once.Do(func() {
		Service = &serviceRegister{
			callList: make(map[string]*callObj),
			rcvList:  make(map[string]interface{}),
		}
		eth := &eth.Eth{}
		Service.Registration(eth)
	})
}

func (a *serviceRegister) Registration(rcv interface{}) error {
	rVal := reflect.ValueOf(rcv)
	rType := rVal.Type()

	if rType.Kind() != reflect.Ptr {
		err := fmt.Errorf("Service Registration want point")
		g.Log().Error(err)
		return err
	}
	rName := rType.Elem().Name()
	pkgName := filepath.Base(rType.Elem().PkgPath())
	if rName == "" || pkgName == "" {
		err := fmt.Errorf("no rName or pkgName?")
		g.Log().Error(err)
		return err
	}

	if _, ok := a.rcvList[rName]; ok {
		err := fmt.Errorf("reRegistration:%s", rName)
		g.Log().Error(err)
		return err
	}
	a.rcvList[rName] = rcv

	// method
	for i := 0; i < rType.NumMethod(); i++ {
		m := rType.Method(i)

		if m.PkgPath != "" {
			g.Log().Warning("Registration Api:", m.Name, ",is not export")
			continue
		}

		mName := formatName(pkgName, m.Name)
		if _, ok := a.callList[mName]; ok {
			g.Log().Error("Registration Api:", mName, ",is existing")
			continue
		}

		call := buildCallObj(m.Func, rVal)
		if call == nil {
			continue
		}
		a.callList[mName] = call
	}

	g.Log().Debug(rName, "Register:", len(a.callList))
	for k, _ := range a.callList {
		g.Log().Debug("method:", k)
	}
	return nil
}

var contextType = reflect.TypeOf((*context.Context)(nil)).Elem()
var errType = reflect.TypeOf((*error)(nil)).Elem()

func formatName(pkgname, name string) string {
	ret := []rune(name)
	if len(ret) > 0 {
		ret[0] = unicode.ToLower(ret[0])
	}
	return pkgname + "_" + string(ret)
}

func (a *serviceRegister) Call(method string, msg *rpc.JsonMessage) (*rpc.JsonMessage, error) {
	call, ok := a.callList[method]
	if !ok {
		return nil, &rpc.MethodNotFoundError{Method: method}
	}

	//build args
	args, err := parseArgs(msg.Params, call.argsType)
	if err != nil {
		return nil, err
	}
	//

	ctx := context.Background()
	ret, err := call.call(ctx, args)
	m := &rpc.JsonMessage{
		JsonRpc: msg.JsonRpc,
		Id:      msg.Id,
	}
	if err != nil {
		return nil, err
	}
	rst, err := json.Marshal(ret)
	if err != nil {
		return nil, err
	}
	m.Result = rst
	return m, nil
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
		if argsType[i].Kind() != reflect.Ptr {
			return nil, fmt.Errorf("missing value for required argument %d", i)
		}
		args[i] = reflect.Zero(argsType[i])
	}
	return args, nil
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
