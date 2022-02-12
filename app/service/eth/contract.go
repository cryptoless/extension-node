package eth

import (
	"context"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

func (e *Eth) GetStorageAt(
	ctx context.Context,
	account common.Address,
	key string,
	blockNrOrHash rpc.BlockNumberOrHash) (hexutil.Bytes, error) {
	return nil, nil
}

type OverrideAccount struct {
	Nonce     *hexutil.Uint64              `json:"nonce"`
	Code      *hexutil.Bytes               `json:"code"`
	Balance   **hexutil.Big                `json:"balance"`
	State     *map[common.Hash]common.Hash `json:"state"`
	StateDiff *map[common.Hash]common.Hash `json:"stateDiff"`
}

func (e *Eth) Call(
	ctx context.Context,
	msg ethereum.CallMsg,
	blockNrOrHash rpc.BlockNumberOrHash,
	overrides *map[common.Address]OverrideAccount) (hexutil.Bytes, error) {

	return nil, nil

}
func (e *Eth) GetProof(
	ctx context.Context,
	address common.Address,
	storageKeys []string,
	blockNrOrHash rpc.BlockNumberOrHash) (*gethclient.AccountResult, error) {
	return nil, nil
}
