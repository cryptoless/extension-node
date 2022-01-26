package model

type GetTokenAllowanceReq struct {
	ExtApiReq
	Params []GetTokenAllowanceParam `json:"params"`
}
type GetTokenAllowanceResp struct {
	ExtApiResp
	Result interface{} `json:"result,omitempty"`
}

///
///
type GetTokenAllowanceParam struct {
	Contract string `json:"contract"`
	Owner    string `json:"owner"`
	Spender  string `json:"spender"`
}
type GetTokenAllowanceResult struct {
	Result string `json:"result"`
}
