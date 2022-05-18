package public

import (
	"golang-test2/models"
	schema_pasien "golang-test2/schemas/pasien"
)

type (
	PublicRoot struct {
		Pasien schema_pasien.Get `json:"pasien"`
		Kamar  []models.Kamars   `json:"kamar"`
	}
)
