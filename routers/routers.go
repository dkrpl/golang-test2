package routers

import (
	"golang-test2/controllers"
	"golang-test2/middleware"

	"github.com/gin-gonic/gin"
)

type Routers struct {
	engine      *gin.Engine
	controllers *controllers.Controllers
	dynamic     *middleware.Dynamic
}
type Router interface {
	Register()
}

func NewRouters(engine *gin.Engine, controller *controllers.Controllers, dynamic *middleware.Dynamic) *Routers {
	return &Routers{
		engine:      engine,
		controllers: controller,
		dynamic:     dynamic,
	}
}

func (this_ *Routers) Register() {
	this_.engine.Use(this_.dynamic.Use())
	this_.registerControllerPasien()
	// this_.registerControllerUser()
	this_.registerControllerKamar()
	// this_.registerSwagger()
	this_.registerPublic()
}
