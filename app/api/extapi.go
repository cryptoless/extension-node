package api

import (
	"extension-node/app/service"
	"extension-node/util/response"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var ExtApi = extApi{}

type extApi struct{}

func (*extApi) Api(r *ghttp.Request) {
	ws, err := r.WebSocket()
	if err != nil {
		// http
		msg := service.ParseMessage(r.GetBody())
		rst := service.Service.HandleMsg(r.GetCtx(), msg)
		response.Response(r, rst)
	} else {
		// ws, until to close
		service.WsCon(r.GetCtx(), ws).Poll()
		g.Log().Debug("close ws.")
	}
}
