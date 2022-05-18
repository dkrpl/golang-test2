package routers

import "golang-test2/config"

func (this_ *Routers) registerPublic() {
	directory := config.GetDir()
	this_.engine.Static("/assets", directory+"/assets/")
	this_.engine.GET("/", this_.controllers.ControllerPublic.Root)
}
