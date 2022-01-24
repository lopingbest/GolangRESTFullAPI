package repository

import (
	"context"
	"database/sql"
	"errors"
	"lopingbest/GolangRESTFullAPI/helper"
	"lopingbest/GolangRESTFullAPI/model/domain"
)

type CategoryRespositoryImplementation struct {
}

func NewCategoryRespositoryImplementation() CategoryRepository {
	return &CategoryRespositoryImplementation{}
}

func (c *CategoryRespositoryImplementation) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "insert into customer(name values (?)"
	result, err := tx.ExecContext(ctx, SQL, category.Name) //category.name merupakan target
	helper.PanicIfError(err)
	id, err := result.LastInsertId() //mendapatkan id terakhir
	helper.PanicIfError(err)

	//set category
	category.Id = int(id)
	return category
}

func (c *CategoryRespositoryImplementation) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "update category set name = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, category.Name, category.Id)
	helper.PanicIfError(err)

	return category
}

func (c *CategoryRespositoryImplementation) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	SQL := "delete from category where id = ?"
	_, err := tx.ExecContext(ctx, SQL, category.Id)
	helper.PanicIfError(err)

}

func (repository *CategoryRespositoryImplementation) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {
	SQL := "select id, name from category where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, categoryId)
	helper.PanicIfError(err)
	defer rows.Close()

	category := domain.Category{}
	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		return category, nil
	} else {
		return category, errors.New("category is not found")
	}
}

func (c *CategoryRespositoryImplementation) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	SQL := "select id, name from category"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)

	//data slice
	var categories []domain.Category

	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		categories = append(categories, category)
	}
	return categories
}
