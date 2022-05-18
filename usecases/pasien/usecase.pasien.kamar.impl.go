package pasien

import (
	schemas "golang-test2/schemas/pasien"

	"gopkg.in/mgo.v2/bson"
)

func (this_ *UsecasePasienImpl) AddKamar(id string, schema schemas.AddKamar) (schemas.AddKamar_Response, error) {

	err := this_.RepositoryPasien.AddKamar(bson.ObjectIdHex(id), schema)
	if err != nil {
		return schemas.AddKamar_Response{
			Message: err.Error(),
			Status:  "error",
		}, err
	} else {
		return schemas.AddKamar_Response{
			Message: "success",
			Status:  "success",
		}, nil
	}
}
