package service

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"authorization_service/pkg/pb"
)

type Service struct {
	repository RepositoryInterface
	signingKey string
	expireTime time.Duration
}

func NewService(repository RepositoryInterface) *Service{
	err := godotenv.Load()
	if err != nil{
		log.Fatal(err)
	}

	signingKey := os.Getenv("SIGNING_KEY")
	if len(signingKey) == 0{
		log.Fatal(errors.New("empty signing key"))
	}
	expirationTime, err := strconv.Atoi(os.Getenv("EXPIRATION_TIME"))
	if err != nil ||expirationTime == 0{
		log.Fatal(errors.New("wrong expiration time"))
	}

	return &Service{
		repository: repository,
		signingKey: signingKey,
		expireTime: time.Duration(expirationTime) * time.Hour,
	}
}

func (s *Service) SignIn(ctx context.Context, req *pb.SignInRequest) (*pb.SignInResponse, error) {
	hashedPassword, _ := hashPassword(req.GetPassword())

	if s.repository.CheckAuthData(ctx, req.GetBusinessId(), hashedPassword) {
		token, err := s.generateToken(req.BusinessId)
		if err != nil {
			return nil, err
		}

		err = s.repository.InitSession(ctx, req.GetBusinessId(), token)
		if err != nil{
			return nil, status.Errorf(codes.Internal, "unable to sign in")
		}

		return &pb.SignInResponse{
			AccessToken: token,
		}, nil
	} else {
		return nil, errors.New("Authorization failed")
	}
}

func (s *Service) AuthorizeUser(ctx context.Context, req *pb.AuthorizationRequest) (*pb.AuthorizationResponse, error) {
	businessID, err := s.parseToken(req.GetToken())
	if err != nil{
		return nil, status.Errorf(codes.InvalidArgument, "invalid access token")
	}

	authStatus, err := s.repository.CheckSession(ctx, businessID, req.GetToken())
	if err != nil{
		return nil, status.Errorf(codes.Aborted, "session error")
	}

	return &pb.AuthorizationResponse{
		AuthStatus: authStatus,
	}, nil
}



func (s *Service) generateToken(business_id string) (string, error) {
	tkn := jwt.NewWithClaims(jwt.SigningMethodES512, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(s.expireTime).Unix(),
		},
		business_id,
	})

	key, err := base64.URLEncoding.DecodeString(s.signingKey)
	if err != nil{
		log.Println(err)
		return "", errors.New("Unable to generate token")
	}

	token, err := tkn.SignedString(key)
	if err != nil{
		log.Println(err)
		return "", err
	}

	return token, err
}

type tokenClaims struct {
	jwt.StandardClaims
	Id     string `json:"id"`
}

func hashPassword(password string) (string, error) {
	hashByte, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return string(hashByte), err
}


func (s *Service) parseToken(accessToken string) (string, error){
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}

		key, err := base64.URLEncoding.DecodeString(s.signingKey)
		if err != nil {
			log.Println(err)
		}

		return key, nil
	})

	if err != nil {
		log.Println("Parse token error: ", err)
		return "", err
	}

	if claims, ok := token.Claims.(*tokenClaims); ok && token.Valid {
		return claims.Id, nil
	} else {
		return "", errors.New("Error invalid access token")
	}
}

type RepositoryInterface interface {
	CheckAuthData(ctx context.Context, business_id, hashed_password string) bool
	InitSession(ctx context.Context, business_id, token string) error   
	CheckSession(ctx context.Context, business_id, token string) (pb.AuthorizationStatus, error)
}
