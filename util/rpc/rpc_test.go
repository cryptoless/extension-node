package rpc

import (
	"encoding/json"
	"testing"

	"github.com/gogf/gf/test/gtest"
)

func Test_Rpc(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		msg := JsonMessage{}
		json.Unmarshal([]byte(`{
			"jsonrpc":"2.0",
			"method":"eth_getBalance",
			"params":[
				"0x04d78999accdb4446763ba2c002cf8d8651643be",
				"0xb71b01"
			],
			"id":1
		}`), &msg)
		t.AssertEQ(msg.Method, "eth_getBalance")
	})
}
