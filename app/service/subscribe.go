package service

import (
	"context"
	"extension-node/app/model"
	"fmt"
	"sync"

	"github.com/gogf/gf/frame/g"
)

type Subscriber struct {
	sync.RWMutex
	ctx      context.Context
	notifier *notifier

	callables map[string]*CallAble
}

func newSubscriber(ctx context.Context, notifier *notifier) *Subscriber {

	sub := &Subscriber{
		ctx:       ctx,
		notifier:  notifier,
		callables: make(map[string]*CallAble, 0),
	}

	go sub.run()
	return sub
}
func (s *Subscriber) Subscribe(msg *model.JsonMessage) *model.JsonMessage {
	//todo: ratelimit
	ca, err := Service.CallAble(s.ctx, msg.Method, msg)
	if err != nil {
		g.Log().Error(err)
		return model.ErrorMessage(err)
	}
	if !ca.obj.isSub {
		err := fmt.Errorf("notifications not supported")
		g.Log().Error(err)
		return model.ErrorMessage(err)

	}
	if err != nil {
		g.Log().Error(err)
		return model.ErrorMessage(err)
	}

	id, err := s.subscribe(ca)
	if err != nil {
		g.Log().Error(err)
		return model.ErrorMessage(err)
	}
	return msg.Response(id)

}
func (s *Subscriber) UnSubscribe(msg *model.JsonMessage) {
	// todo parse id
	s.Lock()
	defer s.Unlock()
	if _, ok := s.callables["id"]; ok {
		delete(s.callables, "id")
	}
}

func (s *Subscriber) run() {
	fmt.Println("Subscriber run")
	for {
		select {
		case <-s.ctx.Done():

			s.Lock()
			s.callables = make(map[string]*CallAble)
			s.Unlock()

		default:
			//todo: notify
			SyncService.WaitBlock()

			s.RLock()
			if len(s.callables) == 0 {
				s.RUnlock()
				continue
			}
			s.RUnlock()

			fmt.Println("Subscriber default")
		}
		///
		clist := make([]*CallAble, 0)

		s.RLock()
		for _, v := range s.callables {
			clist = append(clist, v)
		}
		s.RUnlock()

		//todo
		for _, c := range clist {
			rst, err := c.Call(s.ctx)
			if err != nil {
				g.Log().Error(err)
				// todo:
				s.notifier.Notify(c.msg.ErrorResponse(err), nil)
			} else {
				s.notifier.Notify(c.msg.Response(rst), nil)
			}

		}
	}
}

func (s *Subscriber) subscribe(ca *CallAble) (string, error) {

	s.Lock()
	defer s.Unlock()
	s.callables[ca.Id()] = ca

	return ca.Id(), nil
}
