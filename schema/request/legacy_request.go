package request

import (
	"database/sql"
	"time"

	"github.com/cpartogi/withdrawdeposit/constant"
	"github.com/cpartogi/withdrawdeposit/entity"
)

// UserRegistrationLegacy for
type UserRegistrationLegacy struct {
	Email          string `validate:"required,email" json:"email,omitempty"`
	Msisdn         string `validate:"required,number,startswith=62,min=10" json:"msisdn,omitempty"`
	Password       string `validate:"required,min=6" json:"password,omitempty"`
	Name           string `validate:"required" json:"name,omitempty"`
	ExternalUserID int64  `validate:"required,number" json:"external_user_id,omitempty"`
}

// CreateUserLegacyParams is
func (r *UserRegistrationLegacy) CreateUserLegacyParams() entity.CreateUserLegacyParams {
	timeNow := time.Now().Unix()
	return entity.CreateUserLegacyParams{
		ExternalUserID: sql.NullInt64{Int64: r.ExternalUserID, Valid: true},
		Name:           r.Name,
		Email:          r.Email,
		Msisdn:         sql.NullString{String: r.Msisdn, Valid: true},
		Password:       r.Password,
		Gender:         entity.GenderM,
		UserStatusID:   constant.StatusUnverified,
		UserTypeID:     constant.TypeUser,
		CreatedAt:      sql.NullInt64{Int64: timeNow, Valid: true},
		UpdatedAt:      sql.NullInt64{Int64: timeNow, Valid: true},
	}
}

// CompanyRegistrationLegacy for
type CompanyRegistrationLegacy struct {
	Email          string `validate:"required,email" json:"email,omitempty"`
	Password       string `validate:"required,min=6" json:"password,omitempty"`
	Name           string `validate:"required" json:"name,omitempty"`
	ExternalUserID int64  `validate:"required,number" json:"external_user_id,omitempty"`
}

// CreateUserLegacyParams is
func (r *CompanyRegistrationLegacy) CreateUserLegacyParams() entity.CreateUserLegacyParams {
	timeNow := time.Now().Unix()
	return entity.CreateUserLegacyParams{
		ExternalUserID: sql.NullInt64{Int64: r.ExternalUserID, Valid: true},
		Name:           r.Name,
		Email:          r.Email,
		Password:       r.Password,
		Gender:         entity.GenderM,
		UserStatusID:   constant.StatusUnverified,
		UserTypeID:     constant.TypeCompany,
		CreatedAt:      sql.NullInt64{Int64: timeNow, Valid: true},
		UpdatedAt:      sql.NullInt64{Int64: timeNow, Valid: true},
	}
}
