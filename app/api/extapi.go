package api

import (
	"encoding/json"
	"extension-node/app/service"
	"extension-node/util/response"
	"extension-node/util/rpc"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var ExtApi = extApi{}

type extApi struct{}

func (*extApi) Api(r *ghttp.Request) {
	ws, err := r.WebSocket()
	if err != nil {
		// http
		rst := api(r.GetBody())
		response.Response(r, rst)
	}

	// ws
	for {
		msgType, msg, err := ws.ReadMessage()
		if err != nil {
			g.Log().Error(err)
			msg := rpc.ErrorMessage(err)
			b, err := json.Marshal(msg)
			if err != nil {
				g.Log().Error(err)
			}
			ws.WriteMessage(msgType, b)
			return
		}
		rst := api(msg)
		b, err := json.Marshal(rst)
		if err != nil {
			g.Log().Error(err)
		}
		if err = ws.WriteMessage(msgType, b); err != nil {
			if err != nil {
				g.Log().Error(err)
			}
			return
		}
	}
}

func api(body []byte) *rpc.JsonMessage {
	req := &rpc.JsonMessage{}
	err := json.Unmarshal(body, req)

	if err != nil {
		g.Log().Error(err)
		return rpc.ErrorMessage(err)
	}
	g.Log().Debug("Api req:", req.Method)

	rst, err := service.Service.Call(req.Method, req)
	if err != nil {
		g.Log().Error(err)
		return rpc.ErrorMessage(err)
	}
	//

	return rst
}
