package dao

import "github.com/globalsign/mgo/bson"

type Keccaks struct {
	Id            bson.ObjectId `bson:"_id"`
	BlockNumber   int           `bson:"blocknumber"`
	Key           string        `bson:"key"`
	fourbyteskey  string        `bson:"fourbyteskey"`
	Type          string        `bson:"type"`
	Preimages     []string      `bson:"preimages"`
	PreimagesSize []int         `bson:"preimagessize"`
	LastPreimage  string        `bson:"lastpreimage"`
}
