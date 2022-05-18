package public

import "github.com/gin-gonic/gin"

type ControllerPublic interface {
	Root(*gin.Context)
}
