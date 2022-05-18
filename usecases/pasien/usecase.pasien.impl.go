package pasien

import (
	"errors"
	helpers "golang-test2/helper"
	"golang-test2/models"
	"golang-test2/pkg"
	pasien_repository "golang-test2/repositorys/pasien"
	pasien_schemas "golang-test2/schemas/pasien"
	"mime/multipart"
	"strconv"
	"strings"

	"gopkg.in/mgo.v2/bson"
)

type UsecaseKamarPasienImpl struct {
}
type UsecasePasienImpl struct {
	RepositoryPasien pasien_repository.PasienRepository
	Dynamic          struct {
		Host string
	}
}

func NewUsecasePasienImpl(repo *pasien_repository.PasienRepository) PasienUsecase {
	return &UsecasePasienImpl{
		RepositoryPasien: *repo,
	}
}

func (this_ *UsecasePasienImpl) Create(schemas pasien_schemas.Input, file *multipart.FileHeader) (pasien_response pasien_schemas.Input_Response, err error) {
	if file == nil {
		schemas.Kabupaten = ""
	} else {
		schemas.Kabupaten, err = helpers.UploadWithReplaceV2("pasien", file)
	}
	// schemas.Icon, err = helpers.Upload("pasien", uuid.New(), schemas.Icon)
	if err != nil {
		return pasien_schemas.Input_Response{}, err
	}
	err = this_.RepositoryPasien.Create(&schemas)
	if err != nil {
		return pasien_schemas.Input_Response{
			Message: err.Error(),
			Status:  "error",
		}, err
	} else {
		return pasien_schemas.Input_Response{
			Message: "success",
			Status:  "success",
		}, nil
	}
}

func (this_ *UsecasePasienImpl) List(search, filter, per_page, page_no string) (pasiens []pasien_schemas.List, count int, err error) {
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
				"name": regex_search,
			},
		}
	} else {
		querys = bson.M{}
	}
	users, count := this_.RepositoryPasien.List(querys, per_page_int, page_no_int)
	for _, pasien := range users {
		pasiens = append(pasiens, pasien_schemas.List{
			ID:          pasien.ID.Hex(),
			Nama_pasien: pasien.Nama_pasien,
			Alamat:      pasien.Alamat,
			No_rumah:    pasien.No_rumah,
			Kabupaten:   pasien.Kabupaten,
			Kecamatan:   pasien.Kecamatan,
		})
	}
	if count == 0 {
		if search != "" && filter != "" {
			err = errors.New("search & filtering not found")
		} else if filter != "" {
			err = errors.New("filtering not found")
		} else if search != "" {
			err = errors.New("search  not found")
		} else {
			err = errors.New("staff is empty")
		}
	}
	return
}

func (this_ *UsecasePasienImpl) Get(id string) (pasien pasien_schemas.Get_Response, err error) {
	pasien_result := this_.RepositoryPasien.Get(bson.ObjectIdHex(id))
	pasien = pasien_schemas.Get_Response{
		ID:          pasien_result.ID.Hex(),
		Nama_pasien: pasien_result.Nama_pasien,
		Alamat:      pasien_result.Alamat,
		No_rumah:    pasien_result.No_rumah,
		Kabupaten:   pasien_result.Kabupaten,
		Kecamatan:   pasien_result.Kecamatan,
	}
	if len(pasien_result.Kamars) == 0 {
		pasien.Kamars = make([]models.Kamars, 0)
	} else {
		pasien.Kamars = pasien_result.Kamars
	}
	return
}

func (this_ *UsecasePasienImpl) Delete(id string) (pasien_schemas.Delete_Response, error) {
	err := this_.RepositoryPasien.Delete(bson.ObjectIdHex(id))
	if err != nil {
		return pasien_schemas.Delete_Response{
			Message: err.Error(),
			Status:  "error",
		}, err
	} else {
		return pasien_schemas.Delete_Response{
			Message: "success",
			Status:  "success",
		}, nil
	}
}

func (this_ *UsecasePasienImpl) Edit(id string, schemas pasien_schemas.Edit) (pasien_response pasien_schemas.Edit_Response, err error) {
	city_repo, err := this_.Get(id)
	if err != nil {
		pasien_response = pasien_schemas.Edit_Response{
			Message: err.Error(),
			Status:  "error",
		}
		return
	}
	if strings.Contains(schemas.Kecamatan, ";base64") {
		err = helpers.UploadWithReplace(city_repo.Kecamatan, schemas.Kecamatan)
		if err != nil {
			return pasien_response, err
		}
	}
	err = this_.RepositoryPasien.Edit(bson.ObjectIdHex(id), schemas)
	if err != nil {
		pasien_response = pasien_schemas.Edit_Response{
			Message: err.Error(),
			Status:  "error",
		}
		return
	} else {
		pasien_response = pasien_schemas.Edit_Response{
			Message: "success",
			Status:  "success",
		}
		return
	}
}
