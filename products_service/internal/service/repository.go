package service

import(
	"context"
	"products_service/internal/models"
)

type RepositoryI interface {
	CreateProduct(ctx context.Context, product *models.CreateProduct) (id string, err error)
	UpdateProduct(ctx context.Context, product *models.Product) (status string, err error)
	DeleteProduct(ctx context.Context, id string) (status string, err error)
	GetProducts()
}