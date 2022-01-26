package dao

import (
	"context"
	"extension-node/boot"
	"time"

	"github.com/globalsign/mgo/bson"
)

type ContractSlotInfo struct {
	Id         bson.ObjectId `bson:"_id"`
	Address    string        `bson:"address"`
	Slot       int           `bson:"slot"`
	Name       string        `base:"name"`
	Type       string        `bson:"type"`
	CreateTime time.Time     `bson:"createtime"`
}

func GetContractSlotInfo(addr string) (*ContractSlotInfo, error) {
	rst := ContractSlotInfo{}
	err := boot.DBContractSlotInfo.Find(context.Background(), bson.M{"address": addr}).One(&rst)
	return &rst, err
}
