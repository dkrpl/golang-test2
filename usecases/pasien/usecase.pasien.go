package pasien

import (
	schemas "golang-test2/schemas/pasien"
	"mime/multipart"
)

type PasienUsecase interface {
	Create(schemas schemas.Input, file *multipart.FileHeader) (schemas.Input_Response, error)
	List(string, string, string, string) ([]schemas.List, int, error)
	Edit(string, schemas.Edit) (schemas.Edit_Response, error)
	Get(string) (schemas.Get_Response, error)
	Delete(string) (schemas.Delete_Response, error)
	AddKamar(string, schemas.AddKamar) (schemas.AddKamar_Response, error)
}
