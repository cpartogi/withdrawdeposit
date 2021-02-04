package request

import (
	"database/sql"
	"time"

	"github.com/cpartogi/withdrawdeposit/entity"
)

// UserRegistration for
type UserRegistration struct {
	Email    string `validate:"required,email" json:"email,omitempty"`
	Msisdn   string `validate:"required,number,startswith=62,min=10" json:"msisdn,omitempty"`
	Password string `validate:"required,min=8" json:"password,omitempty"`
	Name     string `validate:"required" json:"name,omitempty"`
}

// CompanyRegistration for
type CompanyRegistration struct {
	Email    string `validate:"required,email" json:"email,omitempty"`
	Password string `validate:"required,min=8" json:"password,omitempty"`
	Name     string `validate:"required" json:"name,omitempty"`
}

// CreateUserParams is
func (r *UserRegistration) CreateUserParams() entity.CreateUserParams {
	timeNow := time.Now().Unix()
	return entity.CreateUserParams{
		Name:         r.Name,
		Email:        r.Email,
		Msisdn:       sql.NullString{String: r.Msisdn, Valid: true},
		Password:     r.Password,
		Gender:       entity.GenderM,
		UserStatusID: 1,
		UserTypeID:   1,
		CreatedAt:    sql.NullInt64{Int64: timeNow, Valid: true},
		UpdatedAt:    sql.NullInt64{Int64: timeNow, Valid: true},
	}
}

// CreateUserParams is
func (r *CompanyRegistration) CreateUserParams() entity.CreateUserParams {
	timeNow := time.Now().Unix()
	return entity.CreateUserParams{
		Name:         r.Name,
		Email:        r.Email,
		Password:     r.Password,
		Gender:       entity.GenderM,
		UserStatusID: 1,
		UserTypeID:   2,
		CreatedAt:    sql.NullInt64{Int64: timeNow, Valid: true},
		UpdatedAt:    sql.NullInt64{Int64: timeNow, Valid: true},
	}
}
