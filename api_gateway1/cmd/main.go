package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"api_gateway1/pkg/pb/auth_service" 
	"api_gateway1/pkg/pb/products_service"
	"api_gateway1/pkg/pb/users_service"
)

func run() error {
	godotenv.Load("services.env")

	var usersServiceAddr = os.Getenv("USERS_SERVICE_ADDR")
	var productsServiceAddr = os.Getenv("PRODUCTS_SERVICE_ADDR")
	var authServiceAddr = os.Getenv("AUTH_SERVIE_ADDR")
	
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()


	mux := runtime.NewServeMux()
	

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	err := users_service.RegisterSellersUsersServiceHandlerFromEndpoint(ctx, mux, usersServiceAddr, opts)
	if err != nil{
		log.Fatal(err)
	}
	err = products_service.RegisterProductsServiceHandlerFromEndpoint(ctx, mux, productsServiceAddr, opts)
	if err != nil{
		log.Fatal(err)
	}
	err = auth_service.RegisterAuthServiceHandlerFromEndpoint(ctx, mux, authServiceAddr, opts)
	if err != nil{
		log.Fatal(err)
	}

	return http.ListenAndServe(":8085", mux)
}

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

