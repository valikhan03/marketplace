// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ProductsServiceClient is the client API for ProductsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductsServiceClient interface {
	CreateProduct(ctx context.Context, in *CreateProductRequest, opts ...grpc.CallOption) (*CreateProductResponse, error)
	UpdateProduct(ctx context.Context, in *UpdateProductRequest, opts ...grpc.CallOption) (*UpdateProductResponse, error)
	DeleteProduct(ctx context.Context, in *DeleteProductRequest, opts ...grpc.CallOption) (*DeleteProductResponse, error)
	GetAllProducts(ctx context.Context, in *GetAllProductsRequest, opts ...grpc.CallOption) (*GetAllProductsResponse, error)
	GetProductByID(ctx context.Context, in *GetProductByIDRequest, opts ...grpc.CallOption) (*GetProductByIDResponse, error)
}

type productsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProductsServiceClient(cc grpc.ClientConnInterface) ProductsServiceClient {
	return &productsServiceClient{cc}
}

func (c *productsServiceClient) CreateProduct(ctx context.Context, in *CreateProductRequest, opts ...grpc.CallOption) (*CreateProductResponse, error) {
	out := new(CreateProductResponse)
	err := c.cc.Invoke(ctx, "/protobuf.ProductsService/CreateProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productsServiceClient) UpdateProduct(ctx context.Context, in *UpdateProductRequest, opts ...grpc.CallOption) (*UpdateProductResponse, error) {
	out := new(UpdateProductResponse)
	err := c.cc.Invoke(ctx, "/protobuf.ProductsService/UpdateProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productsServiceClient) DeleteProduct(ctx context.Context, in *DeleteProductRequest, opts ...grpc.CallOption) (*DeleteProductResponse, error) {
	out := new(DeleteProductResponse)
	err := c.cc.Invoke(ctx, "/protobuf.ProductsService/DeleteProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productsServiceClient) GetAllProducts(ctx context.Context, in *GetAllProductsRequest, opts ...grpc.CallOption) (*GetAllProductsResponse, error) {
	out := new(GetAllProductsResponse)
	err := c.cc.Invoke(ctx, "/protobuf.ProductsService/GetAllProducts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productsServiceClient) GetProductByID(ctx context.Context, in *GetProductByIDRequest, opts ...grpc.CallOption) (*GetProductByIDResponse, error) {
	out := new(GetProductByIDResponse)
	err := c.cc.Invoke(ctx, "/protobuf.ProductsService/GetProductByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductsServiceServer is the server API for ProductsService service.
type ProductsServiceServer interface {
	CreateProduct(context.Context, *CreateProductRequest) (*CreateProductResponse, error)
	UpdateProduct(context.Context, *UpdateProductRequest) (*UpdateProductResponse, error)
	DeleteProduct(context.Context, *DeleteProductRequest) (*DeleteProductResponse, error)
	GetAllProducts(context.Context, *GetAllProductsRequest) (*GetAllProductsResponse, error)
	GetProductByID(context.Context, *GetProductByIDRequest) (*GetProductByIDResponse, error)
}

// UnimplementedProductsServiceServer can be embedded to have forward compatible implementations.
type UnimplementedProductsServiceServer struct {
}

func (*UnimplementedProductsServiceServer) CreateProduct(context.Context, *CreateProductRequest) (*CreateProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProduct not implemented")
}
func (*UnimplementedProductsServiceServer) UpdateProduct(context.Context, *UpdateProductRequest) (*UpdateProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProduct not implemented")
}
func (*UnimplementedProductsServiceServer) DeleteProduct(context.Context, *DeleteProductRequest) (*DeleteProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProduct not implemented")
}
func (*UnimplementedProductsServiceServer) GetAllProducts(context.Context, *GetAllProductsRequest) (*GetAllProductsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllProducts not implemented")
}
func (*UnimplementedProductsServiceServer) GetProductByID(context.Context, *GetProductByIDRequest) (*GetProductByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductByID not implemented")
}

func RegisterProductsServiceServer(s *grpc.Server, srv ProductsServiceServer) {
	s.RegisterService(&_ProductsService_serviceDesc, srv)
}

func _ProductsService_CreateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductsServiceServer).CreateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.ProductsService/CreateProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductsServiceServer).CreateProduct(ctx, req.(*CreateProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductsService_UpdateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductsServiceServer).UpdateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.ProductsService/UpdateProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductsServiceServer).UpdateProduct(ctx, req.(*UpdateProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductsService_DeleteProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductsServiceServer).DeleteProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.ProductsService/DeleteProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductsServiceServer).DeleteProduct(ctx, req.(*DeleteProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductsService_GetAllProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllProductsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductsServiceServer).GetAllProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.ProductsService/GetAllProducts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductsServiceServer).GetAllProducts(ctx, req.(*GetAllProductsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductsService_GetProductByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductsServiceServer).GetProductByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.ProductsService/GetProductByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductsServiceServer).GetProductByID(ctx, req.(*GetProductByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProductsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.ProductsService",
	HandlerType: (*ProductsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProduct",
			Handler:    _ProductsService_CreateProduct_Handler,
		},
		{
			MethodName: "UpdateProduct",
			Handler:    _ProductsService_UpdateProduct_Handler,
		},
		{
			MethodName: "DeleteProduct",
			Handler:    _ProductsService_DeleteProduct_Handler,
		},
		{
			MethodName: "GetAllProducts",
			Handler:    _ProductsService_GetAllProducts_Handler,
		},
		{
			MethodName: "GetProductByID",
			Handler:    _ProductsService_GetProductByID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "product_service.proto",
}