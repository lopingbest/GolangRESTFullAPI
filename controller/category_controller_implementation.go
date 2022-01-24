package controller

import (
	"github.com/julienschmidt/httprouter"
	"lopingbest/GolangRESTFullAPI/helper"
	"lopingbest/GolangRESTFullAPI/model/service"
	"lopingbest/GolangRESTFullAPI/model/web"
	"net/http"
	"strconv"
)

type CategoryControllerImplementation struct {
	//interface tidak memerlukan pointer karena interface
	CategoryService service.CategoryService
}

//Function yang mengekspos CategoryController dan mengimplementasi struct categoryService
func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &CategoryControllerImplementation{
		CategoryService: categoryService,
	}
}

func (controller *CategoryControllerImplementation) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//parsing
	categoryCreateRequest := web.CategoryCreateRequest{}
	helper.ReadFromRequestBody(r, &categoryCreateRequest)

	//panggil service
	categoryResponse := controller.CategoryService.Create(r.Context(), categoryCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "ok",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *CategoryControllerImplementation) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//parsing
	categoryUpdateRequest := web.CategoryUpdateRequest{}
	helper.ReadFromRequestBody(r, &categoryUpdateRequest)

	//konversi jadi string
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	categoryUpdateRequest.Id = id

	//panggil service
	categoryResponse := controller.CategoryService.Update(r.Context(), categoryUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "ok",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *CategoryControllerImplementation) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//konversi jadi string
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	//panggil service
	controller.CategoryService.Delete(r.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "ok",
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *CategoryControllerImplementation) FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//konversi jadi string
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	//panggil service
	categoryResponse := controller.CategoryService.FindById(r.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "ok",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *CategoryControllerImplementation) FIndAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//panggil service
	categoryResponses := controller.CategoryService.FindAll(r.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "ok",
		Data:   categoryResponses,
	}

	helper.WriteToResponseBody(w, webResponse)
}
