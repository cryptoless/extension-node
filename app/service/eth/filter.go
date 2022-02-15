package eth

import (
	"context"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
)

type FilterQuery struct {
	BlockHash *common.Hash     // used by eth_getLogs, return logs only from block with this hash
	FromBlock *hexutil.Big     // beginning of the queried range, nil means genesis block
	ToBlock   *hexutil.Big     // end of the range, nil means latest block
	Addresses []common.Address // restricts matches to events created by specific contracts

	// The Topic list restricts matches to particular event topics. Each event has a list
	// of topics. Topics matches a prefix of that list. An empty element slice matches any
	// topic. Non-empty elements represent an alternative that matches any of the
	// contained topics.
	//
	// Examples:
	// {} or nil          matches any topic list
	// {{A}}              matches topic A in first position
	// {{}, {B}}          matches any topic in first position AND B in second position
	// {{A}, {B}}         matches topic A in first position AND B in second position
	// {{A, B}, {C, D}}   matches topic (A OR B) in first position AND (C OR D) in second position
	Topics [][]common.Hash
}

func (e *Eth) GetLogs(ctx context.Context, q FilterQuery) ([]types.Log, error) {
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
