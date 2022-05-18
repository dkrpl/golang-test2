package kamar

import schemas "golang-test2/schemas/kamar"

type KamarUsecase interface {
	Create(schemas.Input) (schemas.Input_Response, error)
	List(string, string, string, string) ([]schemas.List, int, error)
	Get(string) (schemas.Detail, error)
	Edit(string, schemas.Edit) (schemas.Edit_Response, error)
	Delete(string) (schemas.Delete_Response, error)
}
