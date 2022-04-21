package repository

import(
	"context"

	"products_service/internal/models"

	"github.com/jmoiron/sqlx"
)

type Repository struct{
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository{
	return &Repository{
		db: db,
	}
}

func (r *Repository) CreateProduct(ctx context.Context, product *models.CreateProduct) (id string, err error)
func (r *Repository) UpdateProduct(ctx context.Context, product *models.Product) (status string, err error)
func (r *Repository) DeleteProduct(ctx context.Context, id string) (status string, err error)
func (r *Repository) GetProducts()