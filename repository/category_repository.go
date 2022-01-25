package repository

import (
	"context"
	"database/sql"
	"lopingbest/GolangRESTFullAPI/model/domain"
)

//contract categoryRepository
type CategoryRepository interface {
	//parameter kedua merupakan transaksional (Tx), ketiga merupakan data asli

	//sql dibuat pointer karena merupakan data struct, biar tidak dicopy secara terus menerus
	Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Delete(ctx context.Context, tx *sql.Tx, category domain.Category)
	FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Category
}
