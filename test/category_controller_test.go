package test

import (
	"database/sql"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"io"
	"lopingbest/GolangRESTFullAPI/app"
	"lopingbest/GolangRESTFullAPI/controller"
	"lopingbest/GolangRESTFullAPI/helper"
	"lopingbest/GolangRESTFullAPI/middleware"
	"lopingbest/GolangRESTFullAPI/repository"
	"lopingbest/GolangRESTFullAPI/service"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func setupTestDB() *sql.DB {
	db, err := sql.Open("mysql", "root:123@tcp(localhost:3306)/golang_restful_api_test")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func setupRouter(db *sql.DB) http.Handler {
	validate := validator.New()
	CategoryRepository := repository.NewCategoryRespositoryImplementation()
	categoryservice := service.NewCategoryService(CategoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryservice)
	router := app.NewRouter(categoryController)

	return middleware.NewAuthMiddleware(router)
}

//delete data sebelum running, karena sebelumnya data terus bertambah ketika akan merunning code
func truncateCategory(db *sql.DB) {
	db.Exec("TRUNCATE category")
}

func TestCreateCategorySuccess(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name": "Gadget"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/categories", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-KEY", "SECRET")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	//baca body
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{} //data bisa berubah ubah
	json.Unmarshal(body, &responseBody)

	//fmt.Println(responseBody) //baca responseBody

	//pengecekan mendalam
	assert.Equal(t, 200, int(responseBody["code"].(float64))) //float64 dikonversi menjadi integer
	assert.Equal(t, "ok", responseBody["status"])
	assert.Equal(t, "Gadget", responseBody["data"].(map[string]interface{})["name"])

}

func TestCreateCategoryFailed(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name": ""}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/categories", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-KEY", "SECRET")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	//baca body
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{} //data bisa berubah ubah
	json.Unmarshal(body, &responseBody)

	//fmt.Println(responseBody) //baca responseBody

	//pengecekan mendalam
	assert.Equal(t, 400, int(responseBody["code"].(float64))) //float64 dikonversi menjadi integer
	assert.Equal(t, "BAD REQUEST", responseBody["status"])
}

func TestUpdateCategorySuccess(t *testing.T) {

}

func TestUpdateCategoryFailed(t *testing.T) {

}

func TestGetCategorySuccess(t *testing.T) {

}

func TestGetCategoryFailed(t *testing.T) {

}

func TestDeleteCategorySuccess(t *testing.T) {

}

func TestCategoryFailed(t *testing.T) {

}

func TestLIstCategoriesSuccess(t *testing.T) {

}

func TestUnauthorized(t *testing.T) {

}
