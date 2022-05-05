package main

import(
	"fmt"
	"io/ioutil"
	"os"
	"log"
	"net"

	"gopkg.in/yaml.v2"
	"github.com/jmoiron/sqlx"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"

	"products_service/pkg/repository"
	"products_service/pkg/service"
	"products_service/pkg/pb"
)

func main() {
	products_repository := repository.NewRepository(initDB())
	products_service := service.NewProductsService(products_repository)

	grpcServer := grpc.NewServer()
	pb.RegisterProductsServiceServer(grpcServer, products_service)

	if err := godotenv.Load(); err != nil{
		log.Fatal(err)
	}

	listener, err := net.Listen("tcp", os.Getenv("SERVER_ADDR"))
	if err != nil{
		log.Fatal(err)
	}

	fmt.Println("Server started working..")
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
}

type db_configs struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Database string `yaml:"database"`
	User     string `yaml:"user"`
	SSLMode  string `yaml:"sslmode"`
}

func initDB() *sqlx.DB {
	var configs db_configs
	data, err := ioutil.ReadFile("config/configs.yaml")
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(data, &configs)
	if err != nil {
		log.Fatal(err)
	}
	err = godotenv.Load()
	if err != nil{
		log.Fatal(err)
	}

	data_src := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		configs.Host, configs.Port, configs.User, os.Getenv("DB_PASSWORD"), configs.Database, configs.SSLMode)

	db, err := sqlx.Open("pgx", data_src)
	if err != nil {
		log.Fatal(err)
	}
	return db
}