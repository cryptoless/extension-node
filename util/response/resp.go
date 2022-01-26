package response

import (
	"extension-node/app/model"

	"github.com/gogf/gf/net/ghttp"
)

func Json(r *ghttp.Request, data ...interface{}) {
	responseData := interface{}(nil)
	if len(data) > 0 {
		responseData = data[0]
	}

	r.Response.WriteJson(responseData)
}

func JsonExit(r *ghttp.Request, data ...interface{}) {
	Json(r, data...)
	r.Exit()
}

func ErrExit(r *ghttp.Request, err error) {
	resp := model.ExtApiResp{
		Err: err.Error(),
	}
	Json(r, resp)
	r.Exit()
}
