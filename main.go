package main

import (
	"trainig_api/controller"
	"trainig_api/db"
	"trainig_api/repository"
	"trainig_api/router"
	"trainig_api/usecase"
)

func main() {
	db := db.InitDB()
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)
	e := router.NewRouter(userController)
	e.Logger.Fatal(e.Start(":8080"))
}
