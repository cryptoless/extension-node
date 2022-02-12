package eth

import (
	"context"
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

func (e *Eth) GetBalance(
	ctx context.Context,
	account common.Address,
	blockNrOrHash rpc.BlockNumberOrHash) (*hexutil.Big, error) {
	ret := hexutil.Big(*big.NewInt(1234))
	return &ret, nil
}

func (e *Eth) GetCode(ctx context.Context, account common.Address, blockNumber *big.Int) ([]byte, error) {
	return nil, nil
}

func (e *Eth) Account(ctx context.Context, account common.Address) (bool, error) {
	return true, nil
}

func (e *Eth) ProtocolVersion(ctx context.Context, q ethereum.FilterQuery) (int, error) {
	return 99, nil
}
func (e *Eth) GasPrice(ctx context.Context) (*big.Int, error) {
	return nil, nil
}
func (e *Eth) EstimateGas(ctx context.Context, msg ethereum.CallMsg, blockNrOrHash *rpc.BlockNumberOrHash) (*big.Int, error) {
	return big.NewInt(123), nil
}

///
type feeHistoryResult struct {
	OldestBlock  *hexutil.Big     `json:"oldestBlock"`
	Reward       [][]*hexutil.Big `json:"reward,omitempty"`
	BaseFee      []*hexutil.Big   `json:"baseFeePerGas,omitempty"`
	GasUsedRatio []float64        `json:"gasUsedRatio"`
}

func (e *Eth) FeeHistory(
	ctx context.Context,
	blockCount int,
	lastBlock rpc.BlockNumber,
	rewardPercentiles []float64) (*feeHistoryResult, error) {

	return nil, nil
}

func (e *Eth) MaxPriorityFeePerGas(ctx context.Context) (*hexutil.Big, error) {
	return (*hexutil.Big)(big.NewInt(1234)), nil

}

///
func (e *Eth) ChainId(ctx context.Context) (*hexutil.Big, error) {
	return (*hexutil.Big)(big.NewInt(1234)), nil
}

///
func (e *Eth) Net_Version(ctx context.Context) string {
	return "version"
}

func (e *Eth) Net_Listening(ctx context.Context) bool {
	return true
}
func (e *Eth) GetUncleByBlockNumberAndIndex(
	ctx context.Context,
	blockNr rpc.BlockNumber, index hexutil.Uint) (map[string]interface{}, error) {
	return nil, nil
}

func (e *Eth) GetUncleByBlockHashAndIndex(
	ctx context.Context,
	blockHash common.Hash,
	index hexutil.Uint) (map[string]interface{}, error) {
	return nil, nil
}

func (e *Eth) GetUncleCountByBlockNumber(
	ctx context.Context,
	blockNr rpc.BlockNumber) *hexutil.Uint {
	n := hexutil.Uint(123)
	return &n
}
func (e *Eth) GetUncleCountByBlockHash(
	ctx context.Context,
	blockNr rpc.BlockNumber) *hexutil.Uint {
	n := hexutil.Uint(123)
	return &n
}

func (e *Eth) Syncing_subscription(ctx context.Context) (interface{}, error) {
	return map[string]interface{}{
		"startingBlock": big.NewInt(123),
	}, nil
}

///

func (e *Eth) Web3_ClientVersion(ctx context.Context) (*big.Int, error) {
	return nil, nil
}
func (e *Eth) Web3_Sha3(ctx context.Context, data common.Hash) (*big.Int, error) {
	return nil, nil
}
