package eth

import (
	"context"
	"encoding/json"
	"extension-node/app/model"
	"fmt"
)

func (e *Eth) Subscribe(ctx context.Context, method string, params json.RawMessage) (*model.JsonMessage, error) {
	fmt.Println(method, params)

	return &model.JsonMessage{
		Method: method,
		Params: params,
	}, nil
}
func (e *Eth) Unsubscribe(ctx context.Context, ids string) (string, error) {
	return ids, nil
}
