package model

import (
	"errors"
	"fmt"
	"reflect"
)

type GetTokenBalancesReq struct {
	ExtApiReq
	Params []interface{} `p:"params", v:"reqired"`
}

func (a *GetTokenBalancesReq) Parse() (*GetTokenBalancesParam, error) {
	req := GetTokenBalancesParam{}
	for i, v := range a.Params {
		if i == 0 {
			switch v.(type) {
			case string:
				req.Address, _ = v.(string)
			default:
				return nil, errors.New(fmt.Sprintf("GetTokenBalancesReq:%+v", reflect.TypeOf((v))))
			}
		} else if i == 1 {
			switch v.(type) {
			case []interface{}:
				for _, v := range v.([]interface{}) {
					switch v.(type) {
					case string:
						addr, _ := v.(string)
						req.OneOf.ContractList = append(req.OneOf.ContractList, addr)
					default:
						return nil, errors.New(fmt.Sprintf("GetTokenBalancesReq:%+v", reflect.TypeOf((v))))
					}
				}
			default:
				return nil, errors.New(fmt.Sprintf("GetTokenBalancesReq:%+v", reflect.TypeOf((v))))
			}
		}

	}
	return &req, nil
}

type GetTokenBalancesResp struct {
	ExtApiResp
	GetTokenBalancesResult
}

///
///
type GetTokenBalancesParam struct {
	Address string `json:"omitempty"`
	OneOf   struct {
		ContractList []string `json:"omitempty"`
		DefToken     string   `json:"omitempty"`
	} `json:"omitempty"`
}
type GetTokenBalancesResult struct {
	Address       string          `json:"address"`
	TokenBalances []TokenBalances `json:"tokenBalances"`
}
type TokenBalances struct {
	ContractAddress string `json:"contractAddress"`
	TokenBalance    string `json:"tokenBalance"`
	Error           string `json:"error"`
}
