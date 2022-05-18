package main

import (
	"golang-test2/config"
	"golang-test2/controllers"
	"golang-test2/middleware"
	"golang-test2/repositorys"
	"golang-test2/routers"
	"golang-test2/usecases"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	engine.Use(middleware.CORSMiddleware())
	engine.SetTrustedProxies([]string{"192.168.18.2"})
	conf := config.NewConf(".env")
	database := conf.NewMongoDatabase()

	repo := repositorys.NewRepository(database)

	dynamic := middleware.NewDynamic()

	usecase := usecases.MainUsecase(repo, dynamic)

	controllers := controllers.MainController(usecase)

	router := routers.NewRouters(engine, &controllers, dynamic)
	router.Register()

	engine.Run(":1753")
}

// controller->usecase->repository
