package service

import (
	"extension-node/app/model"
	"fmt"
)

type GetTokenMetadata struct {
}

func (a *GetTokenMetadata) Get(param []string) *model.GetTokenMetadataResult {
	fmt.Printf("GetTokenAllowance:%+v", param)

	return &model.GetTokenMetadataResult{}
}
