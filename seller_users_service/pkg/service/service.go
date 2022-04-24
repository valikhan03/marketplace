package service

import(
	"context"
	"sellers_users_service/pkg/pb"
)

type Service struct{
	repository RepositoryIntercase
}


func (s *Service) SignUp(ctx context.Context, req *pb.SignUpRequest) (*pb.SignUpResponse, error){

}

func (s *Service) SignIn(ctx context.Context, req *pb.SignInRequest) (*pb.SignInResponse, error){

}

func (s *Service) AccoundData(ctx context.Context, req *pb.AccountDataRequest) (*pb.AccountDataResponse, error){

}

func (s *Service) CheckAccess(ctx context.Context, req *pb.CheckAccessRequest) (*pb.CheckAccessResponse, error){

}

func (s *Service) mustEmbedUnimplementedSellersUsersServiceServer()


type RepositoryIntercase interface{
	RegisterCompany(data *pb.SignUpRequest) (status string, err error)
	GetCompanyData(business_id, password string) ()
	GetAccountData()
	
}