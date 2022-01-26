package dao

import (
	"context"
	"extension-node/boot"

	"github.com/globalsign/mgo/bson"
)

type States struct {
	Id               bson.ObjectId `bson:"_id"`
	BlockNumber      int           `bson:"blocknumber"`
	BlockHash        string        `bson:"BlockHash"`
	TransactionHash  string        `bson:"TransactionHash"`
	TransactionIndex int           `bson:"transactionindex"`
	Address          string        `bson:"address"`
	Slot             int           `bson:"slot"`
	Key              string        `bson:"key"`
	ValueBefore      string        `bson:"valuebefore"`
	ValueAfter       string        `bson:"valueafter"`
}

func GetStates(Key string) (*States, error) {
	rst := States{}
	err := boot.DBStates.Find(context.Background(), bson.M{"key": Key}).One(&rst)
	return &rst, err
}
