package service

import (
	"context"
	"encoding/json"
	"extension-node/app/service/eth"
	"extension-node/util/model"
	"testing"

	"github.com/gogf/gf/test/gtest"
)

func Test_Service(t *testing.T) {

	gtest.C(t, func(t *gtest.T) {
		e := &eth.Eth{}
		Service.Registration(e)
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
		ctx := context.Background()
		_, err := Service.CallAble(ctx, "eth_getBalance", &msg)
		t.AssertEQ(err, nil)

		_, err = Service.CallAble(ctx, "eth_blockNumber", &msg)
		t.AssertNE(err, nil)

		_, err = Service.CallAble(ctx, "no_method", &msg)
		t.AssertNE(err, nil)

		json.Unmarshal([]byte(`{
			"jsonrpc":"2.0",
			"method":"eth_panicMethod",
			"params":[
			],
			"id":1
		}`), &msg)
		_, err = Service.CallAble(ctx, "eth_panicMethod", &msg)
		t.AssertNE(err, nil)
	})

	gtest.C(t, func(t *gtest.T) {
		e := eth.Eth{}
		err := Service.Registration(e)
		t.AssertNE(err, nil)
	})
}

func Benchmark_Handle(b *testing.B) {
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

	ctx := context.Background()
	for i := 0; i < b.N; i++ {
		Service.HandleMsg(ctx, &msg)
	}
}
