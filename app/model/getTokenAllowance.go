package model

type GetTokenAllowanceReq struct {
	Params []GetTokenAllowanceParam `json:"params"`
}
type GetTokenAllowanceResp struct {
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
