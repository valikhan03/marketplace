package repository

import(
	"context"

	"sellers_users_service/pkg/pb"
	"sellers_users_service/pkg/models"
)

type Repository struct{

}

func (r *Repository) RegisterCompany(ctx context.Context, data *pb.SignUpRequest) (pb.SignUpStatus,  error) {
	
}

func (r *Repository) GetCompanyInfo(ctx context.Context, id string) (models.CompanyInfo, error){

}