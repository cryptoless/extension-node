package eth

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
)

func (e *Eth) GetLogs(ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error) {
	return nil, nil
}
func (e *Eth) GetFilterChanges(ctx context.Context, filterId string) (*big.Int, error) {
	return nil, nil
}
func (e *Eth) GetFilterLogs(ctx context.Context, filterId string) (*big.Int, error) {
	return nil, nil
}
func (e *Eth) NewBlockFilter(ctx context.Context) (*big.Int, error) {
	return nil, nil
}
func (e *Eth) NewFilter(ctx context.Context, q ethereum.FilterQuery) (*big.Int, error) {
	return nil, nil
}
func (e *Eth) NewPendingTransactionFilter(ctx context.Context) (*big.Int, error) {
	return nil, nil
}

//
///
func (e *Eth) UninstallFilter(ctx context.Context, filterId string) (*big.Int, error) {
	return nil, nil
}
