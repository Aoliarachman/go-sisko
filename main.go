package main

import (
	"github.com/go-playground/validator/v10"
	"go-sisko/app"
	"go-sisko/controller"
	"go-sisko/helper"
	"go-sisko/middleware"
	"go-sisko/repository"
	"go-sisko/service"
	"net/http"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	elementaryRepository := repository.NewElementaryRepository()

	//Memanggil constructor service
	//depedency : repository, db , validate

	elementaryService := service.NewElementaryService(elementaryRepository, db, validate)
	elementaryController := controller.NewElementaryController(elementaryService)
	router := app.NewRouter(elementaryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
