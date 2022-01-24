package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"lopingbest/GolangRESTFullAPI/app"
	"lopingbest/GolangRESTFullAPI/controller"
	"lopingbest/GolangRESTFullAPI/helper"
	"lopingbest/GolangRESTFullAPI/model/repository"
	"lopingbest/GolangRESTFullAPI/model/service"
	"net/http"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	CategoryRepository := repository.NewCategoryRespositoryImplementation()
	categoryservice := service.NewCategoryService(CategoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryservice)

	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
