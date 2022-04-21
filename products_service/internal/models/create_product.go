package models

import(
	"products_service/internal/pb"
)

type CreateProduct struct{
	Title string
	Description string
	Categories []string
	Photos []string
	Price int32
	SellerID string
}

func MarshalCreateProductRequest(p *pb.CreateProductRequest) *CreateProduct {
	return &CreateProduct{
		Title: p.Title,
		Description: p.Description,
		Categories: p.Categories,
		Photos: p.Photos,
		Price: p.Price,
		
	}
} 