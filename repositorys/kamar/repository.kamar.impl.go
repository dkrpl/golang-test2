package kamar

import (
	"golang-test2/models"
	schemas "golang-test2/schemas/kamar"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func NewKamarRepository(db *mgo.Database) KamarRepository {
	return &KamarRepositoryImpl{Collection: db.C("kamar")}
}

type KamarRepositoryImpl struct {
	Collection *mgo.Collection
}

func (this_ *KamarRepositoryImpl) Create(schema schemas.Input) error {
	id := bson.NewObjectId()
	return this_.Collection.Insert(bson.M{
		"_id":        id,
		"nama_kamar": schema.Nama_kamar,
		"desc":       schema.Desc,
	})
}

func (this_ *KamarRepositoryImpl) Edit(id bson.ObjectId, schema schemas.Edit) error {
	err := this_.Collection.Update(bson.M{"_id": id},
		bson.M{
			"$set": bson.M{
				"nama": schema.Nama_kamar,
				"desc": schema.Desc,
			},
		},
	)
	return err
}

func (this_ *KamarRepositoryImpl) List(query map[string]interface{}, per_page, page_no int) ([]models.Kamar, int, error) {
	var feature []models.Kamar
	err := this_.Collection.Find(query).Skip((page_no - 1) * per_page).Limit(per_page).All(&feature)
	count, _ := this_.Collection.Find(query).Count()
	return feature, count, err
}

func (this_ *KamarRepositoryImpl) Get(id bson.ObjectId) models.Kamar {
	var feature models.Kamar
	this_.Collection.Find(bson.M{"_id": id}).One(&feature)
	return feature
}

func (this_ *KamarRepositoryImpl) Delete(id bson.ObjectId) error {
	return this_.Collection.Remove(bson.M{"_id": id})
}
