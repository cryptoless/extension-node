package service

import (
	"extension-node/app/model"
)

type GetTokenMetadata struct {
}

func (a *GetTokenMetadata) Get(param []string) *model.GetTokenMetadataResult {

	return &model.GetTokenMetadataResult{}
}
