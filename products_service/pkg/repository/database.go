package repository

import (
	"context"
	"log"

	"products_service/pkg/models"
	"products_service/pkg/pb"

	"github.com/google/uuid"
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

func (r *Repository) CreateProduct(ctx context.Context, product *models.CreateProductRequest) (id string, err error){
	query := "INSERT INTO products (id, title, categories, description, price, photos, seller_id) VALUES ($1, $2, $3, $4, $5, $6, $7)"
	id = uuid.NewString()
	_, err = r.db.ExecContext(ctx, query, id, product.Title, product.Categories, product.Description, product.Price, product.Photos, product.SellerID)
	if err != nil{
		log.Println(err)
	}
	return
}

func (r *Repository) UpdateProduct(ctx context.Context, product *models.Product) (status pb.UpdateProductStatus, err error){
	query := "UPDATE products SET title=$1, categories=$2, description=$3, price=$4, photos=$5, seller_id=$6 WHERE id=$7"
	_, err = r.db.ExecContext(ctx, query, product.Title, product.Categories, product.Description, product.Price, product.Photos, product.SellerID, product.Id)
	if err != nil{
		log.Println(err)
		status = pb.UpdateProductStatus_UNABLE_TO_UPDATE
	}else{
		status = pb.UpdateProductStatus_UPDATED
	}
	return
}


func (r *Repository) DeleteProduct(ctx context.Context, id string) (status pb.DeleteProductStatus, err error) {
	_, err = r.db.ExecContext(ctx, "DELETE FROM products WHERE id=$1", id)
	if err != nil{
		log.Println(err)
		status = pb.DeleteProductStatus_UNABLE_TO_DELETE
	}else{
		status = pb.DeleteProductStatus_DELETED
	}
	return
}

func (r *Repository) GetProductsBySellerID(ctx context.Context, sellerID string) (products []*models.Product, err error){
	query := "SELECT * FROM products WHERE id=$1"
	err = r.db.SelectContext(ctx, &products, query, sellerID)
	if err != nil{
		log.Println(err)
	}
	return
}

func (r *Repository) GetProductByID(ctx context.Context, id string) (product *models.Product, err error){
	query := "SELECT * FROM products WHERE id=$1 LIMIT "
	err = r.db.SelectContext(ctx, &product, query, id)
	if err != nil{
		log.Println(err)
	}
	return
}