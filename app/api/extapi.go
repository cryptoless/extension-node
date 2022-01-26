package api

import (
	"errors"
	"extension-node/app/model"
	"extension-node/app/service"
	"extension-node/util/response"
	"fmt"

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
	var (
		extApiReq *model.ExtApiReq
	)

	err := r.ParseForm(&extApiReq)
	if err != nil {
		response.ErrExit(r, err)
	}
	g.Log().Debug("Api req:", extApiReq)

	rst := model.ExtApiResp{}
	switch extApiReq.Method {
	case model.GETTOKENALLOWANCE:
		var req *model.GetTokenAllowanceReq
		err = r.ParseForm(&req)
		if err != nil {
			response.ErrExit(r, err)
		}
		rst.Result = (&service.GetTokenAllowance{}).Get(req.Params)
	case model.GETTOKENBALANCES:
		var req *model.GetTokenBalancesReq
		err = r.ParseForm(&req)
		if err != nil {
			response.ErrExit(r, err)
		}
		param, err := req.Parse()
		if err != nil {
			response.ErrExit(r, err)
		}
		rst.Result = (&service.GetTokenBalances{}).Get(param)

	case model.GETTOKENMETADATA:
		var req *model.GetTokenMetadataReq
		err = r.ParseForm(&req)
		if err != nil {
			response.ErrExit(r, err)
		}
		(&service.GetTokenMetadata{}).Get(req.Params)

	default:
		err := errors.New(fmt.Sprint("unSupport method:", extApiReq.Method))
		g.Log().Error(err)
		response.ErrExit(r, err)
	}

	////

	//
	//

	response.JsonExit(r, rst)
}
