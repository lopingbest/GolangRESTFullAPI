package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"lopingbest/GolangRESTFullAPI/helper"
	"lopingbest/GolangRESTFullAPI/model/domain"
	"lopingbest/GolangRESTFullAPI/model/repository"
	"lopingbest/GolangRESTFullAPI/model/web"
)

type CategoryServiceImplemenation struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
	validate           *validator.Validate
}

func (service CategoryServiceImplemenation) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	//validasi sebelum mulai
	err := service.validate.Struct(request)
	helper.PanicIfError(err)

	//transaksi dimulai dari service, untuk menanggulangi jika suatu saat ada kasus yang membutuhkan lebih dari satu repository
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category := domain.Category{
		Name: request.Name,
	}

	//data transaksi dikirim dari service ke repository
	category = service.CategoryRepository.Save(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func (service CategoryServiceImplemenation) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	err := service.validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	//divalidasi dulu apakah ada atau tidak
	category, err := service.CategoryRepository.FindById(ctx, tx, request.Id)
	helper.PanicIfError(err)

	//kalo category sudah ketemu, selanjutnya akan di set
	category.Name = request.Name

	category = service.CategoryRepository.Update(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func (service CategoryServiceImplemenation) Delete(ctx context.Context, categoryId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	//divalidasi dulu apakah ada atau tidak
	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	//kalo tidak ada, maka akan panic
	helper.PanicIfError(err)

	//kalo ada, kemudian akan dihapus
	service.CategoryRepository.Delete(ctx, tx, category)
}

func (service CategoryServiceImplemenation) FindById(ctx context.Context, categoryId int) web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	helper.PanicIfError(err)

	return helper.ToCategoryResponse(category)
}

func (service CategoryServiceImplemenation) FindAll(ctx context.Context) []web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	categories := service.CategoryRepository.FindAll(ctx, tx)

	return helper.ToCategoryResponses(categories)
}
