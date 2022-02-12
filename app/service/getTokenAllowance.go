package service

// legacy
import (
	"extension-node/app/dao"
	"extension-node/app/model"
	"extension-node/util/excrypto"
)

type GetTokenAllowance struct {
}

func (a *GetTokenAllowance) Get(param []model.GetTokenAllowanceParam) *model.GetTokenAllowanceResult {
	p := param[0]
	contractSlotInfo, err := dao.GetContractSlotInfo(p.Contract)
	if err != nil {
		panic(err)
	}
	slot := contractSlotInfo.Slot

	// find by key
	h1 := excrypto.ContractSlotKey(int64(slot), p.Owner)
	key := excrypto.ContractKey(p.Spender, h1.String())
	//
	stat, err := dao.GetStates(key.String())
	if err != nil {
		panic(err)
	}

	return &model.GetTokenAllowanceResult{
		Result: stat.ValueAfter,
	}
}
