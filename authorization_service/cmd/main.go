package main

import(
	"log"
	"io/ioutil"
	"os"
	"fmt"
	"net"

	"gopkg.in/yaml.v2"
	"github.com/joho/godotenv"
	"github.com/jmoiron/sqlx"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc"


	"authorization_service/pkg/repository"
	"authorization_service/pkg/service"
	"authorization_service/pkg/pb"

)

func main() {
	auth_repository := repository.NewRepository(initRedisDB(), initDB())
	auth_service := service.NewService(auth_repository)

	grpcServer := grpc.NewServer()

	pb.RegisterAuthServiceServer(grpcServer, auth_service)

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

func initRedisDB() *redis.Client{
	configs := ReadRedisConfigs()
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", configs.Host, configs.Port),
		Password: "",
		DB: 0,
	})

	return rdb
}

type RedisConfigs struct{
	Host string `yaml:"redis.host"`
	Port string `yaml:"redis.port"`
} 

func ReadRedisConfigs() *RedisConfigs {
	var confs RedisConfigs

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

	return &confs
}

func initDB() *sqlx.DB {
	db, err := sqlx.Open("pgx", ReadPostgresConfigs())
	if err != nil {
		log.Fatal(err)
	}
	return db
}

type PostgresConfigs struct {
	Host    string `yaml:"db.host"`
	Port    string `yaml:"db.port"`
	User    string `yaml:"db.user"`
	DBName  string `yaml:"db.dbname"`
	SSLMode string `yaml:"db.sslmode"`
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