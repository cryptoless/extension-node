package model

var GETTOKENALLOWANCE = "getTokenAllowance"
var GETTOKENBALANCES = "getTokenBalances"
var GETTOKENMETADATA = "getTokenMetadata"

type ExtApiReq struct {
	JsonRpc string `p:"jsonrpc" v:"required"`
	ID      int    `p:"id" v:"required"`
	Method  string `p:"method" v:"required"`
}

type ExtApiResp struct {
	Jsonrpc string      `json:"jsonrpc"`
	ID      int         `json:"id"`
	Err     string      `json:"error,omitempty"`
	Result  interface{} `json:"result,omitempty"`
}
