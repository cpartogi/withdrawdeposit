package request

import (
	"database/sql"

	"github.com/cpartogi/withdrawdeposit/entity"
)

// UserLogin for
type UserLogin struct {
	Msisdn   string `validate:"required,number,startswith=62,min=10" json:"msisdn,omitempty"`
	Password string `validate:"required,min=6" json:"password,omitempty"`
}

// User is
func (r *UserLogin) User() entity.User {
	return entity.User{
		Msisdn:   sql.NullString{String: r.Msisdn, Valid: true},
		Password: r.Password,
	}
}

// OTPLogin for
type OTPLogin struct {
	Msisdn string `validate:"required,number,startswith=62,min=10" json:"msisdn,omitempty"`
	OTP    string `validate:"required,min=4,max=5,number" json:"otp,omitempty"`
}

// User is
func (r *OTPLogin) User() entity.User {
	return entity.User{
		Msisdn: sql.NullString{String: r.Msisdn, Valid: true},
	}
}

// CompanyLogin for
type CompanyLogin struct {
	Email    string `validate:"required,email" json:"email,omitempty"`
	Password string `validate:"required,min=6" json:"password,omitempty"`
}

// User is
func (r *CompanyLogin) User() entity.User {
	return entity.User{
		Email:    r.Email,
		Password: r.Password,
	}
}
