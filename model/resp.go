package model

type Resp struct {
	Jsonrpc string `json:"jsonrpc"`
	Id      int    `json:"id"`
	Result  struct {
		Calls []struct {
			BlockNumber         int    `json:"blockNumber"`
			BlockHash           string `json:"blockHash"`
			TransactionHash     string `json:"transactionHash"`
			TransactionIndex    int    `json:"transactionIndex"`
			Type                string `json:"type"`
			From                string `json:"from"`
			To                  string `json:"to"`
			Value               string `json:"value"`
			Gas                 string `json:"gas"`
			GasUsed             string `json:"gasUsed"`
			Input               string `json:"input"`
			Output              string `json:"output"`
			Status              int    `json:"status"`
			Index               string `json:"index"`
			FromNonce           int    `json:"fromNonce"`
			FromNonceDifference int    `json:"fromNonceDifference"`
			ToNonce             int    `json:"toNonce"`
			ToNonceDifference   int    `json:"toNonceDifference"`
		} `json:"calls"`
		Keccaks []struct {
			BlockNumber   int      `json:"blockNumber"`
			Key           string   `json:"key"`
			FourBytesKey  string   `json:"fourBytesKey"`
			Type          string   `json:"type"`
			Preimages     []string `json:"preimages"`
			PreimagesSize []int    `json:"preimagesSize"`
			LastPreimage  string   `json:"lastPreimage"`
		} `json:"keccaks"`
		States []struct {
			BlockNumber      int    `json:"blockNumber"`
			BlockHash        string `json:"blockHash"`
			TransactionHash  string `json:"transactionHash"`
			TransactionIndex int    `json:"transactionIndex"`
			Address          string `json:"address"`
			Slot             int    `json:"slot"`
			Key              string `json:"key"`
			ValueBefore      string `json:"valueBefore"`
			ValueAfter       string `json:"valueAfter"`
		} `json:"states"`
	} `json:"result"`
}
