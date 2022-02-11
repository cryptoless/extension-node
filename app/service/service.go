package service

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"extension-node/app/service/eth"
	"extension-node/util/model"

	"fmt"
	"io"
	"path/filepath"
	"reflect"
	"strings"
	"sync"
	"unicode"

	"github.com/gogf/gf/frame/g"
)

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

		call := buildCallObj(m.Func, rVal, matchIsSub(m.Name) > 0)
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
	pos := matchIsSub(name)
	if pos > 0 {
		name = name[0:pos]
	}

	ret := []rune(name)
	if len(ret) > 0 {
		ret[0] = unicode.ToLower(ret[0])
	}

	return pkgname + "_" + string(ret)
}
func matchIsSub(name string) int {
	return strings.LastIndex(name, "_subscription")
}
func (a *serviceRegister) CallAble(ctx context.Context, method string, msg *model.JsonMessage) (*CallAble, error) {
	call, ok := a.callList[method]
	if !ok {
		return nil, &model.MethodNotFoundError{Method: method}
	}

	//build args
	// todo: parse panic
	args, err := parseArgs(msg.Params, call.argsType)
	if err != nil {
		return nil, err
	}
	//todo: msg
	callable := call.CallAble(ctx, args)
	callable.msg = msg
	return callable, nil

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

func ParseMessage(body []byte) *model.JsonMessage {
	req := &model.JsonMessage{}
	err := json.Unmarshal(body, req)

	if err != nil {
		g.Log().Error(err)
		return model.ErrorMessage(err)
	}
	g.Log().Debug("Api req:", req.Method)
	return req
}

func (a *serviceRegister) HandleMsg(ctx context.Context, msg *model.JsonMessage) *model.JsonMessage {

	callable, err := Service.CallAble(ctx, msg.Method, msg)
	if err != nil {
		g.Log().Error(err)
		return model.ErrorMessage(err)
	}
	//
	ret, err := callable.Call(ctx)
	if err != nil {
		g.Log().Error(err)
		return model.ErrorMessage(err)
	}
	return msg.Response(ret)
}
