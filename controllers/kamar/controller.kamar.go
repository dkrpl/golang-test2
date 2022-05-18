package kamar

import "github.com/gin-gonic/gin"

type ControllerKamar interface {
	Create(*gin.Context)
	List(*gin.Context)
	Get(*gin.Context)
	Delete(*gin.Context)
	Edit(*gin.Context)
}
