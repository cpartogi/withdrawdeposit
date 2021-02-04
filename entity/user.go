package entity

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

// Gender is
type Gender string

const (
	// GenderM  for Male
	GenderM Gender = "M"
	// GenderF  for Female
	GenderF Gender = "F"
)

// Scan for
func (e *Gender) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Gender(s)
	case string:
		*e = Gender(s)
	default:
		return fmt.Errorf("unsupported scan type for Gender: %T", src)
	}
	return nil
}

// User for
type User struct {
	ID               uuid.UUID      `json:"id"`
	ExternalUserID   sql.NullInt64  `json:"external_user_id"`
	Name             string         `json:"name"`
	Email            string         `json:"email"`
	Msisdn           sql.NullString `json:"msisdn"`
	Password         string         `json:"password"`
	UserStatusID     int32          `json:"user_status_id"`
	ProfileImage     sql.NullString `json:"profile_image"`
	Gender           Gender         `json:"gender"`
	UserTypeID       int64          `json:"user_type_id"`
	IsMsisdnVerified bool           `json:"is_msisdn_verified"`
	IsEmailVerified  bool           `json:"is_email_verified"`
	SubscribeID      sql.NullInt32  `json:"subscribe_id"`
	CreatedAt        sql.NullInt64  `json:"created_at"`
	CreatedBy        uuid.UUID      `json:"created_by"`
	UpdatedAt        sql.NullInt64  `json:"updated_at"`
	UpdatedBy        uuid.UUID      `json:"updated_by"`
	DeletedAt        sql.NullInt64  `json:"deleted_at"`
	DeletedBy        uuid.UUID      `json:"deleted_by"`
}

// CreateUserParams for
type CreateUserParams struct {
	Name         string         `json:"name"`
	Email        string         `json:"email"`
	Msisdn       sql.NullString `json:"msisdn"`
	Password     string         `json:"password"`
	UserStatusID int32          `json:"user_status_id"`
	ProfileImage sql.NullString `json:"profile_image"`
	Gender       Gender         `json:"gender"`
	UserTypeID   int64          `json:"user_type_id"`
	CreatedAt    sql.NullInt64  `json:"created_at"`
	CreatedBy    uuid.UUID      `json:"created_by"`
	UpdatedAt    sql.NullInt64  `json:"updated_at"`
	UpdatedBy    uuid.UUID      `json:"updated_by"`
}

// UpdateUserToVerifiedStatusParams for
type UpdateUserToVerifiedStatusParams struct {
	ID               uuid.UUID `json:"id"`
	IsMsisdnVerified bool      `json:"is_msisdn_verified"`
	UserStatusID     int32     `json:"user_status_id"`
}

// CreateUserLegacyParams for
type CreateUserLegacyParams struct {
	ExternalUserID sql.NullInt64  `json:"external_user_id"`
	Name           string         `json:"name"`
	Email          string         `json:"email"`
	Msisdn         sql.NullString `json:"msisdn"`
	Password       string         `json:"password"`
	UserStatusID   int32          `json:"user_status_id"`
	ProfileImage   sql.NullString `json:"profile_image"`
	Gender         Gender         `json:"gender"`
	UserTypeID     int64          `json:"user_type_id"`
	CreatedAt      sql.NullInt64  `json:"created_at"`
	CreatedBy      uuid.UUID      `json:"created_by"`
	UpdatedAt      sql.NullInt64  `json:"updated_at"`
	UpdatedBy      uuid.UUID      `json:"updated_by"`
}
