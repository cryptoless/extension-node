package eth

import (
	"context"
	"encoding/json"
	"extension-node/util/model"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
)

type Eth struct {
}

func (e *Eth) BlockNumber(ctx context.Context) (uint64, error) {
	return 0, nil
}

func (e *Eth) GetBlockByHash(ctx context.Context, hash common.Hash) (*types.Block, error) {
	return nil, nil
}

func (e *Eth) GetBlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error) {
	return nil, nil
}

///
///

///
///

func (e *Eth) GetBalance(ctx context.Context, account common.Address, blockNrOrHash rpc.BlockNumberOrHash) (*hexutil.Big, error) {
	ret := hexutil.Big(*big.NewInt(1234))
	return &ret, nil
}
func (e *Eth) GetCode(ctx context.Context, account common.Address, blockNumber *big.Int) ([]byte, error) {
	return nil, nil
}

func (e *Eth) Account(ctx context.Context, account common.Address, blockNumber *big.Int) ([]byte, error) {
	return nil, nil
}

func (e *Eth) ProtocolVersion(ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error) {
	return nil, nil
}
func (e *Eth) GasPrice(ctx context.Context) (*big.Int, error) {
	return nil, nil

}
func (e *Eth) EstimateGas(ctx context.Context, msg ethereum.CallMsg) (uint64, error) {
	return 0, nil

}

///
///
func (e *Eth) FeeHistory(ctx context.Context, rang uint64, blockNrOrHash rpc.BlockNumberOrHash, ratio []common.Hash, reward []uint64) (uint64, error) {
	return 0, nil

}
func (e *Eth) MaxPriorityFeePerGas(ctx context.Context) (*big.Int, error) {
	return big.NewInt(1234), nil

}

///
func (e *Eth) ChainId(ctx context.Context) (*big.Int, error) {
	return nil, nil
}

///
func (e *Eth) Net_Version(ctx context.Context) (*big.Int, error) {
	return nil, nil
}

func (e *Eth) Net_Listening(ctx context.Context) (*big.Int, error) {
	return nil, nil
}
func (e *Eth) GetUncleByBlockNumberAndIndex(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash, pos uint64) (*big.Int, error) {
	return nil, nil
}

func (e *Eth) GetUncleByBlockHashAndIndex(ctx context.Context, blockHash common.Hash, pos uint64) (*big.Int, error) {
	return nil, nil
}
func (e *Eth) GetUncleCountByBlockHash(ctx context.Context, blockHash common.Hash) (*big.Int, error) {
	return nil, nil
}
func (e *Eth) GetUncleCountByBlockNumber(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) (*big.Int, error) {
	return nil, nil
}

func (e *Eth) Syncing_subscription(ctx context.Context) (*big.Int, error) {
	return big.NewInt(123), nil
}

///
///
func (e *Eth) Subscribe(ctx context.Context, method string, params json.RawMessage) (*model.JsonMessage, error) {
	fmt.Println(method, params)

	return &model.JsonMessage{
		Method: method,
		Params: params,
	}, nil
}
func (e *Eth) Unsubscribe(ctx context.Context, ids []string) ([]string, error) {
	return ids, nil
}

func (e *Eth) Web3_ClientVersion(ctx context.Context) (*big.Int, error) {
	return nil, nil
}
func (e *Eth) Web3_Sha3(ctx context.Context, data common.Hash) (*big.Int, error) {
	return nil, nil
}
