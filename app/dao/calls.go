package dao

import "github.com/globalsign/mgo/bson"

type Calls struct {
	Id bson.ObjectId `bson:"_id"`

	BlockNumber         int    `bson:"blocknumber"`
	BlockHash           string `bson:"BlockHash"`
	TransactionHash     string `bson:"TransactionHash"`
	TransactionIndex    int    `bson:"transactionindex"`
	Type                string `bson:"type"`
	From                string `bson:"from"`
	To                  string `bson:"to"`
	Value               string `bson:"value"`
	Gas                 string `bson:"gas"`
	GasUsed             string `bson:"gasused"`
	Input               string `bson:"input"`
	Output              string `bson:"output"`
	Status              int    `bson:"status"`
	Index               string `bson:"index"`
	FromnOnce           int    `bson:"fromnonce"`
	FromnonceDifference int    `bson:"fromnoncedifference"`
	Tononce             int    `bson:"tononce"`
	TononceDifference   int    `bson:"tononcedifference"`
}

func Get() *Calls {

	return nil
}
