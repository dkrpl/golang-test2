package public

import schema "golang-test2/schemas/public"

type PublicUsecase interface {
	Root(id string) (schema.PublicRoot, error)
}
