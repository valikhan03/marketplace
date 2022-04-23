package service

import (
	"context"
	"errors"

	"products_service/internal/models"
	"products_service/internal/pb"
)

type ProductsService struct {
	repository RepositoryInterface
}

func NewProductsService(r RepositoryInterface) *ProductsService {
	return &ProductsService{
		repository: r,
	}
}

func (s *ProductsService) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	id, err := s.repository.CreateProduct(ctx, models.MarshalCreateProductRequest(req))
	if err != nil {
		return nil, errors.New("DB ERROR")
	}
	return &pb.CreateProductResponse{
		Id: id,
	}, nil
}

func (s *ProductsService) UpdateProduct(ctx context.Context, req *pb.UpdateProductRequest) (*pb.UpdateProductResponse, error) {
	status, err := s.repository.UpdateProduct(ctx, models.MarshalProduct(req.Product))
	if err != nil {
		return nil, errors.New("DB ERROR")
	}
	return &pb.UpdateProductResponse{Status: status}, nil
}

func (s *ProductsService) DeleteProduct(ctx context.Context, req *pb.DeleteProductRequest) (*pb.DeleteProductResponse, error) {
	status, err := s.repository.DeleteProduct(ctx, req.ProductId)
	if err != nil {
		return nil, errors.New("DB ERROR")
	}
	return &pb.DeleteProductResponse{Status: status}, nil
}

func (s *ProductsService) GetAllProducts(ctx context.Context, req *pb.GetAllProductsRequest) (*pb.GetAllProductsResponse, error) {
	prods, err := s.repository.GetProductsBySellerID(ctx, req.GetSellerId())
	if err != nil {
		return nil, errors.New("DB ERROR")
	}

	var products []*pb.Product

	for _, p := range prods {
		go func(p models.Product) {
			products = append(products, models.UnmarshalProduct(&p))
		}(p)
	}

	return &pb.GetAllProductsResponse{Products: products}, err
}

func (s *ProductsService) GetProductByID(ctx context.Context, req *pb.GetProductByIDRequest) (*pb.GetProductByIDResponse, error) {
	product, err := s.repository.GetProductByID(ctx, req.GetId())
	if err != nil {
		return nil, errors.New("DB ERROR")
	}

	if product != nil {
		return &pb.GetProductByIDResponse{Product: models.UnmarshalProduct(product)}, nil
	} else {
		return nil, errors.New("EMPTY RESPONSE FROM DB")
	}
}

func (s *ProductsService) mustEmbedUnimplementedProductsServiceServer() {}

type RepositoryInterface interface {
	CreateProduct(ctx context.Context, product *models.CreateProductRequest) (id string, err error)
	UpdateProduct(ctx context.Context, product *models.Product) (status string, err error)
	DeleteProduct(ctx context.Context, id string) (status string, err error)
	GetProductsBySellerID(ctx context.Context, sellerID string) (products []models.Product, err error)
	GetProductByID(ctx context.Context, id string) (product *models.Product, err error)
}
