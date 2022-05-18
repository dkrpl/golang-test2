package kamar

import (
	"golang-test2/models"
	schemas "golang-test2/schemas/kamar"

	"gopkg.in/mgo.v2/bson"
)

type KamarRepository interface {
	Create(schemas.Input) error
	Edit(bson.ObjectId, schemas.Edit) error
	List(map[string]interface{}, int, int) ([]models.Kamar, int, error)
	Get(bson.ObjectId) models.Kamar
	Delete(bson.ObjectId) error
}
