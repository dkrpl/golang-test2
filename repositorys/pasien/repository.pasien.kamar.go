package pasien

import (
	schemas "golang-test2/schemas/pasien"

	"gopkg.in/mgo.v2/bson"
)

func (this_ *PasienRepositoryImpl) AddKamar(id bson.ObjectId, schema schemas.AddKamar) error {
	return this_.Collection.Update(bson.M{"_id": id},
		bson.M{
			"$Set": bson.M{
				"kamars": bson.M{
					schema.NamaKamar: bson.M{
						"nama": schema.Nama,
						"path": schema.Path,
					},
				},
			},
		},
	)
}
