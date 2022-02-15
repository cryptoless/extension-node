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

func ErrorResponse(r *ghttp.Request, err model.Error) {
	rsp := (&model.JsonMessage{}).ErrorResponse(err)
	JsonExit(r, rsp)
}
func Response(r *ghttp.Request, msg *model.JsonMessage) {
	JsonExit(r, msg)
}
