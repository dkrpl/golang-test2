package models

import "gopkg.in/mgo.v2/bson"

type (
	Pasien struct {
		ID          bson.ObjectId `bson:"_id,omitempty"`
		Nama_pasien string        `bson:"nama_pasien"`
		Alamat      string        `bson:"alamat"`
		No_rumah    int           `json:"no_rumah"`
		Kecamatan   string        `bson:"kecamatan"`
		Kabupaten   string        `bson:"kabupaten"`
		Kamars      Kamars        `bson:"kamars"`
	}
	Kamars map[string]interface{}
)
