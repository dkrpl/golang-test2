package kamar

import (
	"errors"
	"golang-test2/pkg"
	"golang-test2/repositorys"
	kamar "golang-test2/repositorys/kamar"
	pasien "golang-test2/repositorys/pasien"
	schemas "golang-test2/schemas/kamar"
	"strconv"

	"gopkg.in/mgo.v2/bson"
)

func NewUsecaseKamar(repo *repositorys.Repositorys) KamarUsecase {
	return &UsecaseKamarImpl{
		RepositoryPasien: *repo.RepositoryPasien,
		RepositoryKamar:  *repo.RepositoryKamar,
	}
}

type UsecaseKamarImpl struct {
	RepositoryPasien pasien.PasienRepository
	RepositoryKamar  kamar.KamarRepository
}

func (this_ *UsecaseKamarImpl) Create(schema schemas.Input) (kamar_response schemas.Input_Response, err error) {
	err = this_.RepositoryKamar.Create(schema)
	if err != nil {
		return schemas.Input_Response{
			Message: err.Error(),
			Status:  "error",
		}, err
	} else {
		return schemas.Input_Response{
			Message: "success",
			Status:  "success",
		}, nil
	}
}

func (this_ *UsecaseKamarImpl) List(search, filter, per_page, page_no string) (kamars_response []schemas.List, count int, err error) {
	per_page_int, _ := strconv.Atoi(per_page)
	page_no_int, _ := strconv.Atoi(page_no)

	if per_page_int == 0 {
		per_page_int = 5
	}
	if page_no_int == 0 {
		page_no_int = 1
	}
	querys := make(map[string]interface{})
	if filter != "" && search != "" {
		regex_search := bson.RegEx{Pattern: search, Options: "i"}
		querys = pkg.ParseFilter(filter)
		querys["$or"] = []interface{}{
			bson.M{
				"name": regex_search,
			},
		}
	} else if filter != "" {
		querys = pkg.ParseFilter(filter)
	} else if search != "" {
		regex_search := bson.RegEx{Pattern: search, Options: "i"}
		querys["$or"] = []interface{}{
			bson.M{
				"nama": regex_search,
			},
		}
	} else {
		querys = bson.M{}
	}
	kamars, count, _ := this_.RepositoryKamar.List(querys, per_page_int, page_no_int)
	for _, kamar := range kamars {
		kamars_response = append(kamars_response, schemas.List{
			ID:         kamar.ID.Hex(),
			Nama_kamar: kamar.Nama_kamar,
			Desc:       kamar.Desc,
		})
	}
	return
}

func (this_ *UsecaseKamarImpl) Get(id string) (kamar_response schemas.Detail, err error) {
	kamar := this_.RepositoryKamar.Get(bson.ObjectIdHex(id))
	kamar_response = schemas.Detail{
		ID:         kamar.ID.Hex(),
		Nama_kamar: kamar.Nama_kamar,
		Desc:       kamar.Desc,
	}
	return
}

func (this_ *UsecaseKamarImpl) Edit(id string, schema schemas.Edit) (edit_response schemas.Edit_Response, err error) {
	err = this_.RepositoryKamar.Edit(bson.ObjectIdHex(id), schema)
	if err != nil {
		edit_response = schemas.Edit_Response{
			Message: err.Error(),
			Status:  "error",
		}
		return
	} else {
		edit_response = schemas.Edit_Response{
			Message: "success",
			Status:  "success",
		}
		return
	}
}

func (this_ *UsecaseKamarImpl) Delete(id string) (delete_response schemas.Delete_Response, err error) {
	kamar := this_.RepositoryKamar.Get(bson.ObjectIdHex(id))
	pasien, _ := this_.RepositoryPasien.GetByKamar(kamar.Nama_kamar)
	if len(pasien) > 0 {
		delete_response = schemas.Delete_Response{
			Message: "you cannot delete this kamar, because this kamar is used in several pasiens",
			Status:  "error",
		}
		err = errors.New(delete_response.Message)
		return
	}
	err = this_.RepositoryKamar.Delete(bson.ObjectIdHex(id))
	if err != nil {
		return schemas.Delete_Response{
			Message: err.Error(),
			Status:  "error",
		}, err
	} else {
		return schemas.Delete_Response{
			Message: "success",
			Status:  "success",
		}, nil
	}
}
