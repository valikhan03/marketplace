package repository

import (
	"context"
	"errors"
	"log"

	"github.com/jmoiron/sqlx"

	"sellers_users_service/pkg/models"
	"sellers_users_service/pkg/pb"
)

type Repository struct{
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository{
	return &Repository{
		db: db,
	}
}

type existance_result struct{
	exists bool
}

func (r *Repository) RegisterCompany(ctx context.Context, data *pb.SignUpRequest) (pb.SignUpStatus,  error) {
	var ex_res existance_result
	err := r.db.Get(&ex_res, "SELECT exists(select 1 FROM companies WHERE id=$1) AS exists", data.BusinessId)
	if err != nil{
		log.Printf("RegisterCompany(): error reading from db: %x\n", err)
		return -1, errors.New("error reading from db")
	}

	//need to check if ex_res is nil-value

	if ex_res.exists{
		return pb.SignUpStatus_COMPANY_ALREADY_EXISTS, nil
	}else{
		query := "INSERT INTO companies (name, type, business_id, phone_nums, email, password) VALUES ($1, $2, $3, $4, $5, $6)"
		_ = r.db.MustExecContext(ctx, query, data.CompanyName, data.CompanyType, data.BusinessId, data.PhoneNums, data.Email, data.Password)
	
		return pb.SignUpStatus_SIGNED_UP, nil
	}
}

func (r *Repository) GetCompanyInfo(ctx context.Context, id string) (*models.CompanyInfo, error){
	var companyInfo models.CompanyInfo
	err := r.db.Get(&companyInfo, "SELECT * FROM companies WHERE business_id=$1", id)
	if err != nil{
		log.Printf("GetCompanyInfo(): error reading from db: %x\n", err)
		return nil, errors.New("error reading from db")
	}

	return &companyInfo, nil
}