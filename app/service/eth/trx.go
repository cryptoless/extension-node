package eth

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func (e *Eth) GetTransactionByHash(ctx context.Context, hash common.Hash) (tx *types.Transaction, isPending bool, err error) {
	return nil, false, nil
}
func (e *Eth) GetTransactionCount(ctx context.Context, blockHash common.Hash) (uint, error) {
	return 0, nil
}
func (e *Eth) GetTransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	return nil, nil
}

func (e *Eth) GetBlockTransactionCountByHash(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	return nil, nil
}
func (e *Eth) GetBlockTransactionCountByNumber(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	return nil, nil
}
func (e *Eth) GetTransactionByBlockHashAndIndex(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	return nil, nil
}
func (e *Eth) GetTransactionByBlockNumberAndIndex(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	return nil, nil
}
func (e *Eth) SendRawTransaction(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	return nil, nil
}
