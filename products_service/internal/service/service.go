package service

import (
	"context"
	"errors"

	"products_service/internal/models"
	"products_service/internal/pb"
)

type ProductsService struct{
	repository RepositoryI
}

func NewProductsService(r RepositoryI) *ProductsService{
	return &ProductsService{
		repository: r,
	}
}

func (s *ProductsService) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error){
	id, err := s.repository.CreateProduct(ctx, models.MarshalCreateProductRequest(req))
	if err != nil{
		return nil, errors.New("DB ERROR")
	}
	return &pb.CreateProductResponse{
		Id: id,
	}, nil
}

func (s *ProductsService) UpdateProduct(ctx context.Context, req *pb.UpdateProductRequest) (*pb.UpdateProductResponse, error){
	status, err := s.repository.UpdateProduct(ctx, models.MarshalProduct(req.Product))
	if err != nil{
		return nil, errors.New("DB ERROR")
	}
	return &pb.UpdateProductResponse{Status: status}, nil
}

func (s *ProductsService) DeleteProduct(ctx context.Context, req *pb.DeleteProductRequest) (*pb.DeleteProductResponse, error){
	status, err := s.repository.DeleteProduct(ctx, req.ProductId)
	if err != nil{
		return nil, errors.New("DB ERROR")
	}
	return &pb.DeleteProductResponse{Status: status}, nil
}

func (s *ProductsService) GetAllProducts(context.Context, *pb.GetAllProductsRequest) (*pb.GetAllProductsResponse, error){
	return nil, nil
}