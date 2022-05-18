package routers

func (this_ *Routers) registerControllerKamar() {
	router := this_.engine.Group("/kamar")
	router.GET("/list", this_.controllers.ControllerKamar.List)
	// router.Use(middleware.Auth())
	router.POST("/create", this_.controllers.ControllerKamar.Create)
	router.GET("/get/:id", this_.controllers.ControllerKamar.Get)
	router.PUT("/edit/:id", this_.controllers.ControllerKamar.Edit)
	router.DELETE("/delete/:id", this_.controllers.ControllerKamar.Delete)
}
