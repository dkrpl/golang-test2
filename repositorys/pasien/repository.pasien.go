package pasien

import (
	"golang-test2/models"
	schema "golang-test2/schemas/pasien"

	"gopkg.in/mgo.v2/bson"
)

type PasienRepository interface {
	Create(schemas *schema.Input) error
	List(map[string]interface{}, int, int) ([]models.Pasien, int)
	Edit(bson.ObjectId, schema.Edit) error
	Get(bson.ObjectId) models.Pasien
	Delete(bson.ObjectId) error
	GetByNo_rumah(string) models.Pasien
	GetByKamar(string) ([]models.Pasien, error)
	AddKamar(bson.ObjectId, schema.AddKamar) error
}
