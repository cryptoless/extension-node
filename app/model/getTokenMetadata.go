package model

type GetTokenMetadataReq struct {
	ExtApiReq
	Params []string `json:"params"`
}

type GetTokenMetadataResp struct {
	ExtApiResp
	Result GetTokenMetadataResult `json:"result"`
}

///
///
type GetTokenMetadataParam struct {
	Address string
}
type GetTokenMetadataResult struct {
	Logo     string `json:"logo"`
	Symbol   string `json:"symbol"`
	Decimals int    `json:"decimals"`
	Name     string `json:"name"`
}
