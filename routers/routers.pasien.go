package routers

func (this_ *Routers) registerControllerPasien() {
	router := this_.engine.Group("/pasien")
	router.GET("/list", this_.controllers.ControllerPasien.List)
	// router.Use(middleware.Auth())
	router.POST("/create", this_.controllers.ControllerPasien.Create)
	router.GET("/get/:id", this_.controllers.ControllerPasien.Get)
	router.PUT("/edit/:id", this_.controllers.ControllerPasien.Edit)
	router.DELETE("/delete/:id", this_.controllers.ControllerPasien.Delete)

	sub1 := router.Group("/kamar")
	sub1.POST("/add/:id", this_.controllers.ControllerPasien.AddKamar)
}
