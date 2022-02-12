package service

import (
	"context"
	"extension-node/app/model"
	"extension-node/app/service/eth"

	"fmt"
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

func ServiceInit() {
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

	return call.CallAble(ctx, msg)
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
