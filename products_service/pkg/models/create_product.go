package models

import(
	"products_service/pkg/pb"
)

type CreateProductRequest struct{
	Title string
	Description string
	Categories []string
	Photos []string
	Price int32
	SellerID string
}

func MarshalCreateProductRequest(p *pb.CreateProductRequest) *CreateProductRequest {
	return &CreateProductRequest{
		Title: p.Title,
		Description: p.Description,
		Categories: p.Categories,
		Photos: p.Photos,
		Price: p.Price,
		
	}
} 