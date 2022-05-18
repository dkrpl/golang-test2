package public

import (
	"golang-test2/middleware"
	"golang-test2/models"
	"golang-test2/repositorys"
	"golang-test2/repositorys/pasien"
	schema_pasien "golang-test2/schemas/pasien"
	schemas "golang-test2/schemas/public"
)

func NewPublicUsecase(repositorys *repositorys.Repositorys, dynamic *middleware.Dynamic) PublicUsecase {
	return &PublicUsecaseImpl{
		RepositoryPasien: *repositorys.RepositoryPasien,

		Dynamic: dynamic,
	}
}

type PublicUsecaseImpl struct {
	RepositoryPasien pasien.PasienRepository
	Dynamic          *middleware.Dynamic
}

func (this_ *PublicUsecaseImpl) Root(id string) (schema schemas.PublicRoot, err error) {
	if this_.Dynamic.Host == "localhost" || this_.Dynamic.Host == "127.0.0.1" {
		schema = schemas.PublicRoot{}
		schema.Kamar = make([]models.Kamars, 0)
		return

	}
	pasien := this_.RepositoryPasien.GetByNo_rumah(this_.Dynamic.Host)
	schema = schemas.PublicRoot{
		Pasien: schema_pasien.Get{
			ID:          pasien.ID.Hex(),
			Nama_pasien: pasien.Nama_pasien,
			Alamat:      pasien.Alamat,
			No_rumah:    pasien.No_rumah,
			Kecamatan:   pasien.Kecamatan,
			Kabupaten:   pasien.Kecamatan,
		},
	}
	if len(pasien.Kamars) == 0 {
		schema.Kamar = make([]models.Kamars, 0)
	} else {
		schema.Kamar = pasien.Kamars
	}
	return
}
