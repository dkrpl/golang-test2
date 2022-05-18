package pasien

import "github.com/gin-gonic/gin"

type ControllerPasien interface {
	Create(*gin.Context)
	List(*gin.Context)
	Get(*gin.Context)
	Delete(*gin.Context)
	Edit(*gin.Context)

	ControllerKamar
}
type ControllerKamar interface {
	AddKamar(*gin.Context)
}
