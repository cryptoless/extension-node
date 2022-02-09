package service

import (
	"context"
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

type Standard struct {
}

func (s *Standard) ChainID(ctx context.Context) (*big.Int, error) {
	return nil, nil
}

func (s *Standard) BlockByHash(ctx context.Context, hash common.Hash) (*types.Block, error) {
	return nil, nil
}

func (s *Standard) BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error) {
	return nil, nil
}

func (s *Standard) BlockNumber(ctx context.Context) (uint64, error) {
	return 0, nil
}

// HeaderByHash returns the block header with the given hash.
func (s *Standard) HeaderByHash(ctx context.Context, hash common.Hash) (*types.Header, error) {
	return nil, nil
}

// HeaderByNumber returns a block header from the current canonical chain. If number is
// nil, the latest known header is returned.
func (s *Standard) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	return nil, nil
}

type rpcTransaction struct {
	tx *types.Transaction
	txExtraInfo
}

type txExtraInfo struct {
	BlockNumber *string         `json:"blockNumber,omitempty"`
	BlockHash   *common.Hash    `json:"blockHash,omitempty"`
	From        *common.Address `json:"from,omitempty"`
}

func (tx *rpcTransaction) UnmarshalJSON(msg []byte) error {
	if err := json.Unmarshal(msg, &tx.tx); err != nil {
		return err
	}
	return json.Unmarshal(msg, &tx.txExtraInfo)
}

// TransactionByHash returns the transaction with the given hash.
func (s *Standard) TransactionByHash(ctx context.Context, hash common.Hash) (tx *types.Transaction, isPending bool, err error) {
	return nil, false, nil
}

// TransactionSender returns the sender address of the given transaction. The transaction
// must be known to the remote node and included in the blockchain at the given block and
// index. The sender is the one derived by the protocol at the time of inclusion.
//
// There is a fast-path for transactions retrieved by TransactionByHash and
// TransactionInBlock. Getting their sender address can be done without an RPC interaction.
func (s *Standard) TransactionSender(ctx context.Context, tx *types.Transaction, block common.Hash, index uint) (common.Address, error) {
	return common.HexToAddress("0x1"), nil
}

// TransactionCount returns the total number of transactions in the given block.
func (s *Standard) TransactionCount(ctx context.Context, blockHash common.Hash) (uint, error) {
	return 0, nil
}

// TransactionInBlock returns a single transaction at index in the given block.
func (s *Standard) TransactionInBlock(ctx context.Context, blockHash common.Hash, index uint) (*types.Transaction, error) {
	return nil, nil
}

// TransactionReceipt returns the receipt of a transaction by transaction hash.
// Note that the receipt is not available for pending transactions.
func (s *Standard) TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	return nil, nil
}

// SyncProgress retrieves the current progress of the sync algorithm. If there's
// no sync currently running, it returns nil.
func (s *Standard) SyncProgress(ctx context.Context) (*ethereum.SyncProgress, error) {
	return nil, nil
}

// SubscribeNewHead subscribes to notifications about the current blockchain head
// on the given channel.
// func (s *Standard) SubscribeNewHead(ctx context.Context, ch chan<- *types.Header) (ethereum.Subscription, error) {
// 	return nil, nil
// }

// State Access

// NetworkID returns the network ID (also known as the chain ID) for this chain.
func (s *Standard) NetworkID(ctx context.Context) (*big.Int, error) {
	return nil, nil
}

// StorageAt returns the value of key in the contract storage of the given account.
// The block number can be nil, in which case the value is taken from the latest known block.
func (s *Standard) StorageAt(ctx context.Context, account common.Address, key common.Hash, blockNumber *big.Int) ([]byte, error) {
	return nil, nil
}

// CodeAt returns the contract code of the given account.
// The block number can be nil, in which case the code is taken from the latest known block.
func (s *Standard) CodeAt(ctx context.Context, account common.Address, blockNumber *big.Int) ([]byte, error) {
	return nil, nil
}

// NonceAt returns the account nonce of the given account.
// The block number can be nil, in which case the nonce is taken from the latest known block.
func (s *Standard) NonceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (uint64, error) {
	return 0, nil
}

// Filters

// FilterLogs executes a filter query.
func (s *Standard) FilterLogs(ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error) {
	return nil, nil
}

// SubscribeFilterLogs subscribes to the results of a streaming filter query.
// func (s *Standard) SubscribeFilterLogs(ctx context.Context, q ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error) {
// 	arg, err := toFilterArg(q)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return ec.c.EthSubscribe(ctx, ch, "logs", arg)
// }

// func toFilterArg(q ethereum.FilterQuery) (interface{}, error) {
// 	arg := map[string]interface{}{
// 		"address": q.Addresses,
// 		"topics":  q.Topics,
// 	}
// 	if q.BlockHash != nil {
// 		arg["blockHash"] = *q.BlockHash
// 		if q.FromBlock != nil || q.ToBlock != nil {
// 			return nil, fmt.Errorf("cannot specify both BlockHash and FromBlock/ToBlock")
// 		}
// 	} else {
// 		if q.FromBlock == nil {
// 			arg["fromBlock"] = "0x0"
// 		} else {
// 			arg["fromBlock"] = toBlockNumArg(q.FromBlock)
// 		}
// 		arg["toBlock"] = toBlockNumArg(q.ToBlock)
// 	}
// 	return arg, nil
// }

// Pending State

// PendingBalanceAt returns the wei balance of the given account in the pending state.
func (s *Standard) PendingBalanceAt(ctx context.Context, account common.Address) (*big.Int, error) {
	return nil, nil
}

// PendingStorageAt returns the value of key in the contract storage of the given account in the pending state.
func (s *Standard) PendingStorageAt(ctx context.Context, account common.Address, key common.Hash) ([]byte, error) {
	return nil, nil

}

// PendingCodeAt returns the contract code of the given account in the pending state.
func (s *Standard) PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error) {
	return nil, nil

}

// PendingNonceAt returns the account nonce of the given account in the pending state.
// This is the nonce that should be used for the next transaction.
func (s *Standard) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
	return 0, nil

}

// PendingTransactionCount returns the total number of transactions in the pending state.
func (s *Standard) PendingTransactionCount(ctx context.Context) (uint, error) {
	return 0, nil

}

func (s *Standard) CallContract(ctx context.Context, msg ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	return nil, nil

}

// PendingCallContract executes a message call transaction using the EVM.
// The state seen by the contract call is the pending state.
func (s *Standard) PendingCallContract(ctx context.Context, msg ethereum.CallMsg) ([]byte, error) {
	return nil, nil

}

// SuggestGasPrice retrieves the currently suggested gas price to allow a timely
// execution of a transaction.
func (s *Standard) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	return nil, nil

}

// SuggestGasTipCap retrieves the currently suggested gas tip cap after 1559 to
// allow a timely execution of a transaction.
func (s *Standard) SuggestGasTipCap(ctx context.Context) (*big.Int, error) {
	return nil, nil

}

// EstimateGas tries to estimate the gas needed to execute a specific transaction based on
// the current pending state of the backend blockchain. There is no guarantee that this is
// the true gas limit requirement as other transactions may be added or removed by miners,
// but it should provide a basis for setting a reasonable default.
func (s *Standard) EstimateGas(ctx context.Context, msg ethereum.CallMsg) (uint64, error) {
	return 0, nil

}

// SendTransaction injects a signed transaction into the pending pool for execution.
//
// If the transaction was a contract creation use the TransactionReceipt method to get the
// contract address after the transaction has been mined.
func (s *Standard) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	return nil
}

func (s *Standard) GetBalance(ctx context.Context, account common.Address, blockNrOrHash rpc.BlockNumberOrHash) (*hexutil.Big, error) {
	ret := hexutil.Big(*big.NewInt(1234))
	return &ret, nil
}
func (s *Standard) GetProof(ctx context.Context, address common.Address, storageKeys []string, blockNrOrHash rpc.BlockNumberOrHash) (*gethclient.AccountResult, error) {
	return nil, nil
}
func (s *Standard) GetBlockByNumber(ctx context.Context, number rpc.BlockNumber, fullTx bool) (map[string]interface{}, error) {
	return nil, nil
}
