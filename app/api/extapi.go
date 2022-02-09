package api

import (
	"encoding/json"
	"extension-node/app/model"
	"extension-node/app/service"
	"extension-node/util/response"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var ExtApi = extApi{}

type extApi struct{}

func Json(r *ghttp.Request, data ...interface{}) {
	responseData := interface{}(nil)
	if len(data) > 0 {
		responseData = data[0]
	}

	r.Response.WriteJson(responseData)
}

func (*extApi) Api(r *ghttp.Request) {

	req := model.JsonMessage{}
	err := json.Unmarshal(r.GetBody(), &req)
	if err != nil {
		response.ErrExit(r, err)
	}
	g.Log().Debug("Api req:", req.Method)

	rst, err := service.Service.Call(req.Method, &req)
	if err != nil {
		response.ErrExit(r, err)
	}
	//

	response.JsonExit(r, rst)
}
