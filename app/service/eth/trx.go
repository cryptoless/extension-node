package eth

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
)

func (e *Eth) GetTransactionByHash(ctx context.Context, hash common.Hash) (tx *types.Transaction, isPending bool, err error) {
	return nil, false, nil
}
func (e *Eth) GetTransactionCount(
	ctx context.Context,
	address common.Address,
	blockNrOrHash rpc.BlockNumberOrHash) (uint, error) {
	return 0, nil
}
func (e *Eth) GetTransactionReceipt(ctx context.Context, blockHash common.Hash) (*types.Receipt, error) {
	return nil, nil
}

func (e *Eth) GetBlockTransactionCountByHash(ctx context.Context, blockHash common.Hash) (uint, error) {
	return 0, nil
}
func (e *Eth) GetBlockTransactionCountByNumber(ctx context.Context, blockNr rpc.BlockNumber) (uint, error) {
	return 1, nil
}
func (e *Eth) GetTransactionByBlockHashAndIndex(ctx context.Context, blockHash common.Hash, index hexutil.Uint) (*types.Transaction, error) {
	return nil, nil
}
func (e *Eth) GetTransactionByBlockNumberAndIndex(ctx context.Context, blockHash common.Hash, index hexutil.Uint) (*types.Transaction, error) {
	return nil, nil
}
func (e *Eth) SendRawTransaction(ctx context.Context, input hexutil.Bytes) (common.Hash, error) {
	return common.HexToHash("0x123"), nil
}
