package service

import (
	"context"

	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
	
	"products_service/pkg/models"
	"products_service/pkg/pb"
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
		return nil, status.Error(codes.Internal, "unable to create product")
	}
	return &pb.CreateProductResponse{
		Id: id,
	}, nil
}

func (s *ProductsService) UpdateProduct(ctx context.Context, req *pb.UpdateProductRequest) (*pb.UpdateProductResponse, error) {
	upd_status, err := s.repository.UpdateProduct(ctx, models.MarshalProduct(req.Product))
	if err != nil {
		return nil, status.Error(codes.Internal, "unable to update product")
	}
	return &pb.UpdateProductResponse{UpdateStatus: upd_status}, nil
}

func (s *ProductsService) DeleteProduct(ctx context.Context, req *pb.DeleteProductRequest) (*pb.DeleteProductResponse, error) {
	del_status, err := s.repository.DeleteProduct(ctx, req.ProductId)
	if err != nil {
		return nil, status.Error(codes.Internal, "unable to delete product")
	}
	return &pb.DeleteProductResponse{DeleteStatus: del_status}, nil
}

func (s *ProductsService) GetAllProducts(ctx context.Context, req *pb.GetAllProductsRequest) (*pb.GetAllProductsResponse, error) {
	prods, err := s.repository.GetProductsBySellerID(ctx, req.GetSellerId())
	if err != nil {
		return nil, status.Error(codes.Internal, "unable to find products")
	}

	var products []*pb.Product

	for _, p := range prods {
		go func(p *models.Product) {
			products = append(products, models.UnmarshalProduct(p))
		}(p)
	}

	return &pb.GetAllProductsResponse{Products: products}, nil
}

func (s *ProductsService) GetProductByID(ctx context.Context, req *pb.GetProductByIDRequest) (*pb.GetProductByIDResponse, error) {
	product, err := s.repository.GetProductByID(ctx, req.GetProductId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unable to find product with id = %s", req.GetProductId())
	}

	if product != nil {
		return &pb.GetProductByIDResponse{Product: models.UnmarshalProduct(product)}, nil
	} else {
		return nil, status.Errorf(codes.NotFound, "unable to find product with id = %s", req.GetProductId())
	}
}

func (s *ProductsService) mustEmbedUnimplementedProductsServiceServer() {}

type RepositoryInterface interface {
	CreateProduct(ctx context.Context, product *models.CreateProductRequest) (id string, err error)
	UpdateProduct(ctx context.Context, product *models.Product) (status pb.UpdateProductStatus, err error)
	DeleteProduct(ctx context.Context, id string) (status pb.DeleteProductStatus, err error)
	GetProductsBySellerID(ctx context.Context, sellerID string) (products []*models.Product, err error)
	GetProductByID(ctx context.Context, id string) (product *models.Product, err error)
}
