package service

import (
	"context"
	"encoding/json"
	"extension-node/app/model"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gorilla/websocket"
)

type WsService struct {
	ctx      context.Context
	ws       *ghttp.WebSocket
	sub      *Subscriber
	notifier *notifier
}

func WsCon(ctx context.Context, ws *ghttp.WebSocket) *WsService {
	s := &WsService{
		ctx: ctx,
		ws:  ws,
	}
	s.notifier = NewNotifier(ctx, s.Send)
	s.sub = newSubscriber(ctx, s.notifier)
	return s
}
func (a *WsService) Send(msg *model.JsonMessage) {
	b, err := json.Marshal(msg)
	if err != nil {
		g.Log().Error(err)
	}
	// todo: default errMsg
	err = a.ws.WriteMessage(websocket.TextMessage, b)
	if err != nil {
		g.Log().Error(err)
	}
}

func (a *WsService) SendOrErr(rst interface{}, err model.Error) {
	if err != nil {
		a.Send(model.ErrorMessage(err))
	} else {
		msg := model.ParseMessageData(rst)
		a.Send(msg)
	}
}

func (a *WsService) Poll() {
	for {
		msgType, body, err := a.ws.ReadMessage()
		if err != nil {
			g.Log().Error(err)
			body := model.ErrorMessage(model.ParseError(err.Error()))
			b, err := json.Marshal(body)
			if err != nil {
				g.Log().Error(err)
			}
			a.ws.WriteMessage(msgType, b)
			return
		}

		switch msgType {
		case websocket.TextMessage:
		case -1, websocket.CloseMessage:
			g.Log().Debug("ws closed")
			return
		default:
			g.Log().Warning("UnProcess wsType:", msgType)
			return
		}

		// todo: default errMsg
		msg := model.ParseMessage(body)
		if msg.IsError() {
			b, err := json.Marshal(msg)
			if err != nil {
				g.Log().Error(err)
			}
			err = a.ws.WriteMessage(msgType, b)
			g.Log().Error(err)
		}
		//
		//
		if msg.IsSubscribe() {
			rst := Service.HandleMsg(a.ctx, msg)
			rst = msg.SubscriptionResult(rst)
			msg := a.sub.Subscribe(rst)
			a.Send(msg)
		} else if msg.IsUnsubscribe() {
			msg := a.sub.UnSubscribe(msg)
			a.Send(msg)
		} else {
			rst := Service.HandleMsg(a.ctx, msg)
			a.Send(rst)
		}
	}
}
