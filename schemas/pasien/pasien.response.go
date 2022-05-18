package pasien

import (
	"golang-test2/models"
)

type (
	List struct {
		ID          string `bson:"_id,omitempty"`
		Nama_pasien string `bson:"nama_pasien"`
		Alamat      string `bson:"alamat"`
		No_rumah    int    `json:"no_rumah"`
		Kecamatan   string `bson:"kecamatan"`
		Kabupaten   string `bson:"kabupaten"`
	}
	Get struct {
		ID          string `bson:"_id,omitempty"`
		Nama_pasien string `bson:"nama_pasien"`
		Alamat      string `bson:"alamat"`
		No_rumah    int    `json:"no_rumah"`
		Kecamatan   string `bson:"kecamatan"`
		Kabupaten   string `bson:"kabupaten"`
	}
	Input_Response struct {
		Message string `json:"message"`
		Status  string `json:"status"`
	}
	Edit_Response struct {
		Message string `json:"message"`
		Status  string `json:"status"`
	}
	Delete_Response struct {
		Message string `json:"message"`
		Status  string `json:"status"`
	}
	Get_Response struct {
		ID          string          `bson:"_id,omitempty"`
		Nama_pasien string          `bson:"nama_pasien"`
		Alamat      string          `bson:"alamat"`
		No_rumah    int             `json:"no_rumah"`
		Kecamatan   string          `bson:"kecamatan"`
		Kabupaten   string          `bson:"kabupaten"`
		Kamars      []models.Kamars `json:"kamars"`
	}
	AddKamar_Response struct {
		Message string `json:"message"`
		Status  string `json:"status"`
	}
)
