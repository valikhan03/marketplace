package service

import (
	"time"
	"context"
	"log"

	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"sellers_users_service/pkg/models"
	"sellers_users_service/pkg/pb"
)

type Service struct {
	repository RepositoryInterface
	signingKey string
	expireTime time.Duration
}

func (s *Service) SignUp(ctx context.Context, req *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	hash, err := hashPassword(req.Password)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal error")
	}
	req.Password = hash

	signupStatus, err := s.repository.RegisterCompany(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	res := &pb.SignUpResponse{
		SignUpStatus: signupStatus,
	}

	switch signupStatus {
	case pb.SignUpStatus_SIGNED_UP:
		res.Info = "Your business account successfully registered!"
	case pb.SignUpStatus_WRONG_DATA:
		res.Info = "Your data is not correct. Please, check it and try again."
	case pb.SignUpStatus_COMPANY_ALREADY_EXISTS:
		res.Info = "Your business is already registered. Please, sign in."
	default:
		res.Info = "Unknown error happened..."
	}

	return res, nil
}

func (s *Service) CompanyInfo(ctx context.Context, req *pb.CompanyInfoRequest) (*pb.CompanyInfoResponse, error) {
	info, err := s.repository.GetCompanyInfo(ctx, req.GetBusinessId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	return &pb.CompanyInfoResponse{
		CompanyName: info.CompanyName,
		CompanyType: info.CompanyType,
		BusinessId:  info.BusinessId,
		PhoneNums:   info.PhoneNums,
		Email:       info.Email,
	}, nil
}


func hashPassword(password string) (string, error) {
	hashByte, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return string(hashByte), err
}



func (s *Service) mustEmbedUnimplementedSellersUsersServiceServer() {}

type RepositoryInterface interface {
	RegisterCompany(ctx context.Context, data *pb.SignUpRequest) (pb.SignUpStatus, error)
	GetCompanyInfo(ctx context.Context, id string) (models.CompanyInfo, error)
}
