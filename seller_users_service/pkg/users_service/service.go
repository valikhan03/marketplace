package service

import (
	"time"
	"context"
	"errors"
	"sellers_users_service/pkg/models"
	"sellers_users_service/pkg/pb"
)

type Service struct {
	repository RepositoryIntercase
	signingKey string
	expireTime time.Duration
}

func (s *Service) SignUp(ctx context.Context, req *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	hash, err := hashPassword(req.Password)
	if err != nil {
		return nil, errors.New("unable to provide security")
	}
	req.Password = hash

	status, err := s.repository.RegisterCompany(ctx, req)
	if err != nil {
		return nil, errors.New("DB ERROR")
	}

	res := &pb.SignUpResponse{
		Status: status,
	}

	switch status {
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

func (s *Service) SignIn(ctx context.Context, req *pb.SignInRequest) (*pb.SignInResponse, error) {
	hashedPassword, _ := hashPassword(req.GetPassword())

	if s.repository.CheckAuthData(ctx, req.GetBusinessId(), hashedPassword) {
		token, err := s.auth.GenerateToken(req.BusinessId)
		if err != nil {
			return nil, err
		}

		return &pb.SignInResponse{
			AccessToken: token,
		}, nil
	} else {
		return nil, errors.New("Authorization failed")
	}
}

func (s *Service) CompanyInfo(ctx context.Context, req *pb.CompanyInfoRequest) (*pb.CompanyInfoResponse, error) {
	info, err := s.repository.GetCompanyInfo(ctx, req.GetBusinessId())
	if err != nil {
		return nil, errors.New("DB ERROR")
	}

	return &pb.CompanyInfoResponse{
		CompanyName: info.CompanyName,
		CompanyType: info.CompanyType,
		BusinessId:  info.BusinessId,
		PhoneNums:   info.PhoneNums,
		Email:       info.Email,
	}, nil
}



func (s *Service) mustEmbedUnimplementedSellersUsersServiceServer() {}

type RepositoryIntercase interface {
	RegisterCompany(ctx context.Context, data *pb.SignUpRequest) (status pb.SignUpStatus, err error)
	GetCompanyInfo(ctx context.Context, id string) (models.CompanyInfo, error)
}
