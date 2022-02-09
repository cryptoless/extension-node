package service

import (
	"encoding/json"
	"extension-node/app/model"
	"testing"

	"github.com/gogf/gf/test/gtest"
)

func Test_Service(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		Standard := &Standard{}
		Service.Registration(Standard)

		msg := model.JsonMessage{}
		json.Unmarshal([]byte(`{
			"jsonrpc":"2.0",
			"method":"eth_getBalance",
			"params":[
				"0x04d78999accdb4446763ba2c002cf8d8651643be",
				"0xb71b01"
			],
			"id":1
		}`), &msg)
		_, err := Service.Call("eth_getBalance", &msg)
		t.AssertEQ(err, nil)
	})
}
