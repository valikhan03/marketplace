package repository

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"

	"authorization_service/pkg/pb"
)

type Repository struct {
	redisClient *redis.Client
	postgres    *sqlx.DB
}

func NewRepository(rc *redis.Client, pc *sqlx.DB) *Repository {
	return &Repository{
		redisClient: rc,
		postgres: pc,
	}
}

type authData struct{
	business_id string
	hashed_password string 
}

func (r *Repository) CheckAuthData(ctx context.Context, business_id, hashed_password string) bool {
	var authdata authData
	err := r.postgres.Get(&authdata, "SELECT business_id, hashed_password FROM users WHERE business_id=$1", business_id)
	if err != nil{
		return false
	}
	if business_id == authdata.business_id && hashed_password == authdata.hashed_password{
		return true
	}else{
		return false
	}
}


func (r *Repository) InitSession(ctx context.Context, business_id, token string) error {
	res := r.redisClient.Append(ctx, business_id, token)
	if res.Err() != nil{
		log.Println(res.Err())
		return res.Err()
	}
	return nil
}

func (r *Repository) CheckSession(ctx context.Context, business_id, token string) (pb.AuthorizationStatus, error) {
	res, err := r.redisClient.Get(ctx, business_id).Result()
	if err != nil{
		return -1, err
	}
	if res == token{
		return pb.AuthorizationStatus_VERIFIED, nil
	}else{
		return pb.AuthorizationStatus_FORBIDDEN, nil
	}
}

