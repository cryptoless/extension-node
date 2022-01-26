package boot

import (
	"context"
	"extension-node/config"

	"github.com/qiniu/qmgo"
)

var Mgo *qmgo.Database = nil

var DBCalls *qmgo.Collection = nil
var DBContractSlotInfo *qmgo.Collection = nil
var DBStates *qmgo.Collection = nil

func MongoInit() {
	//qorm
	ctx := context.Background()
	client, err := qmgo.NewClient(ctx, &qmgo.Config{Uri: config.MongoCfg.Url,
		Auth: &qmgo.Credential{
			Username: config.MongoCfg.User,
			Password: config.MongoCfg.Pass,
		}})
	if err != nil {
		panic(err)
	}
	Mgo = client.Database(config.MongoCfg.Db)
	DBCalls = Mgo.Collection("calls")
	DBContractSlotInfo = Mgo.Collection("contractslotinfo")
	DBStates = Mgo.Collection("states")
}
