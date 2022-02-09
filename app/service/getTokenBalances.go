package service

import (
	"extension-node/app/dao"
	"extension-node/app/model"
	"extension-node/util/excrypto"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type GetTokenBalances struct {
}

func (a *GetTokenBalances) Get(param *model.GetTokenBalancesParam) *model.GetTokenBalancesResult {
	rst := model.GetTokenBalancesResult{}
	rst.Address = param.Address

	for _, v := range param.OneOf.ContractList {
		contractSlotInfo, err := dao.GetContractSlotInfo(v)
		if err != nil {
			rst.TokenBalances = append(rst.TokenBalances, model.TokenBalances{
				ContractAddress: v,
				TokenBalance:    common.Bytes2Hex(common.LeftPadBytes(big.NewInt(0).Bytes(), 32)),
				Error:           err.Error(),
			})
			continue
		}
		//calculation key
		key := excrypto.ContractSlotKey(int64(contractSlotInfo.Slot), param.Address)
		//
		stat, err := dao.GetStates(key.String())
		if err != nil {
			rst.TokenBalances = append(rst.TokenBalances, model.TokenBalances{
				ContractAddress: v,
				TokenBalance:    common.Bytes2Hex(common.LeftPadBytes(big.NewInt(0).Bytes(), 32)),
				Error:           err.Error(),
			})
			continue
		}
		rst.TokenBalances = append(rst.TokenBalances, model.TokenBalances{
			ContractAddress: v,
			TokenBalance:    stat.ValueAfter,
			Error:           err.Error(),
		})
	}
	return &rst
}
