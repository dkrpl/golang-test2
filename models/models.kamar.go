package models

import "gopkg.in/mgo.v2/bson"

type Kamar struct {
	ID         bson.ObjectId `bson:"_id,omitempty"`
	Nama_kamar string        `bson:"name_kamar"`
	Desc       string        `json:"desc"`
}
