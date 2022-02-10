package service

import (
	"encoding/json"
	"extension-node/app/service/eth"
	"extension-node/util/rpc"
	"testing"

	"github.com/gogf/gf/test/gtest"
)

func Test_Service(t *testing.T) {

	gtest.C(t, func(t *gtest.T) {
		e := &eth.Eth{}
		Service.Registration(e)
		msg := rpc.JsonMessage{}
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

		_, err = Service.Call("eth_blockNumber", &msg)
		t.AssertNE(err, nil)

		_, err = Service.Call("no_method", &msg)
		t.AssertNE(err, nil)

		json.Unmarshal([]byte(`{
			"jsonrpc":"2.0",
			"method":"eth_panicMethod",
			"params":[
			],
			"id":1
		}`), &msg)
		_, err = Service.Call("eth_panicMethod", &msg)
		t.AssertNE(err, nil)
	})

	gtest.C(t, func(t *gtest.T) {
		e := eth.Eth{}
		err := Service.Registration(e)
		t.AssertNE(err, nil)
	})
}
