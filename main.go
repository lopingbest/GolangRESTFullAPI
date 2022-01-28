package main

import (
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"lopingbest/GolangRESTFullAPI/app"
	"lopingbest/GolangRESTFullAPI/controller"
	"lopingbest/GolangRESTFullAPI/helper"
	"lopingbest/GolangRESTFullAPI/middleware"
	"lopingbest/GolangRESTFullAPI/repository"
	"lopingbest/GolangRESTFullAPI/service"
	"net/http"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	CategoryRepository := repository.NewCategoryRespositoryImplementation()
	categoryservice := service.NewCategoryService(CategoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryservice)

	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
