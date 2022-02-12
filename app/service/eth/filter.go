package eth

import (
	"context"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
)

func (e *Eth) GetLogs(ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error) {
	return nil, nil
}
func (e *Eth) GetFilterChanges(ctx context.Context, filterId string) (interface{}, error) {
	return nil, nil
}
func (e *Eth) GetFilterLogs(ctx context.Context, filterId string) ([]*types.Log, error) {
	return nil, nil
}
func (e *Eth) NewBlockFilter() (string, error) {
	return "id", nil
}
func (e *Eth) NewFilter(ctx context.Context, q ethereum.FilterQuery) (string, error) {
	return "id", nil
}
func (e *Eth) NewPendingTransactionFilter(ctx context.Context) (string, error) {
	return "id", nil
}

//
///
func (e *Eth) UninstallFilter(ctx context.Context, filterId string) (bool, error) {
	return true, nil
}
