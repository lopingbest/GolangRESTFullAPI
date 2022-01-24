package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type CategoryController interface {
	//function mengikuti API yang akan dibuat. Parameter mengikuti http handler. Params diikutkan, menyesuaikan kontrak
	Create(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Update(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	FIndAll(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}
