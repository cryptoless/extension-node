package model

type Req struct {
	JsonRpc string   `json:"jsonrpc"`
	Method  string   `json:"method"`
	Params  []string `json:"params"`
	Id      int      `json:"id"`
}

func NewReq() *Req {
	return &Req{
		JsonRpc: "2.0",
		Method:  "eth2_getBlockContractStates",
		Id:      1,
	}
}
