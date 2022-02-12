package service

import (
	"context"
	"encoding/json"
	"extension-node/app/model"
	"sync"
	"time"
)

type notifier struct {
	sendFn func(*model.JsonMessage)
}

func NewNotifier(ctx context.Context, sendFn func(*model.JsonMessage)) *notifier {
	return &notifier{
		sendFn: sendFn,
	}
}
func (a *notifier) Notify(rst *model.JsonMessage, err error) error {
	if err != nil {
		return a.send(model.ErrorMessage(err))
	}
	_, retErr := json.Marshal(rst)
	if retErr != nil {
		return a.send(model.ErrorMessage(err))
	} else {
		return a.send(rst)
	}
}

func (a *notifier) send(rst *model.JsonMessage) error {
	// todo: buffer
	a.sendFn(rst)
	return nil
}

type syncService struct{}

var SyncService *syncService

var syncServiceOnce sync.Once

func init() {
	syncServiceOnce.Do(func() {
		SyncService = &syncService{}
	})
}

func (a *syncService) WaitBlock() {
	//todo: wait blocksyncing
	time.Sleep(100000000)
}
