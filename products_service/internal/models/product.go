package models

import (
	"products_service/internal/pb"
)

type Product struct {
	Id          string
	Title       string
	Description string
	Categories  []string
	Photos      []string
	Price       int32
	SellerID    string
}

func MarshalProduct(p *pb.Product) *Product {
	return &Product{
		Id:          p.Id,
		Title:       p.Title,
		Description: p.Description,
		Categories:  p.Categories,
		Photos:      p.Photos,
		Price:       p.Price,
		SellerID:    p.SellerId,
	}
}

func UnmarshalProduct(p *Product) *pb.Product {
	return &pb.Product{
		Id:          p.Id,
		Title:       p.Title,
		Description: p.Description,
		Categories:  p.Categories,
		Photos:      p.Photos,
		Price:       p.Price,
		SellerId:    p.SellerID,
	}
}
