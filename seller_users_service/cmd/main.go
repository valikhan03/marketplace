package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"gopkg.in/yaml.v2"

	"sellers_users_service/pkg/pb"
	repository "sellers_users_service/pkg/repository"
	service "sellers_users_service/pkg/users_service"
)

func main() {
	users_repository := repository.NewRepository(initDB())
	users_service := service.NewService(users_repository)

	if err := godotenv.Load(); err != nil{
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterSellersUsersServiceServer(grpcServer, users_service)
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

type PostgresConfigs struct {
	Host    string `yaml:"db.host"`
	Port    string `yaml:"db.port"`
	User    string `yaml:"db.user"`
	DBName  string `yaml:"db.dbname"`
	SSLMode string `yaml:"db.sslmode"`
}

func initDB() *sqlx.DB {
	db, err := sqlx.Open("pgx", ReadPostgresConfigs())
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func ReadPostgresConfigs() string {
	var confs PostgresConfigs

	file, err := os.Open("config/configs.yaml")
	if err != nil {
		log.Fatal(err)
	}

	byteConfigData, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(byteConfigData, &confs)
	if err != nil {
		log.Fatal(err)
	}

	godotenv.Load()
	password := os.Getenv("POSTGRES_PASSWORD")

	conn_str := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s password=%s", confs.Host, confs.Port, confs.User,  confs.DBName, confs.SSLMode, password)

	return conn_str
}