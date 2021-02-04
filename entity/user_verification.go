package entity

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"

	"github.com/google/uuid"
)

// UserVerification for
type UserVerification struct {
	ID               int64            `json:"id"`
	UserID           uuid.UUID        `json:"user_id"`
	VerificationCode VerificationCode `json:"verification_code"`
	EmailToken       EmailToken       `json:"email_token"`
	ResetPassword    ResetPassword    `json:"reset_password"`
	CreatedAt        sql.NullInt64    `json:"created_at"`
	CreatedBy        uuid.UUID        `json:"created_by"`
	UpdatedAt        sql.NullInt64    `json:"updated_at"`
	UpdatedBy        uuid.UUID        `json:"updated_by"`
	DeletedAt        sql.NullInt64    `json:"deleted_at"`
	DeletedBy        uuid.UUID        `json:"deleted_by"`
}

// VerificationCode for
type VerificationCode struct {
	Otp             string        `json:"otp"`
	RequestCount    int           `json:"request_count"`
	FailedCount     int           `json:"failed_count"`
	HoldUntil       sql.NullInt64 `json:"hold_until"`
	LastRequestedAt sql.NullInt64 `json:"last_requested_at"`
	TokenCreatedAt  sql.NullInt64 `json:"token_created_at"`
}

// EmailToken for
type EmailToken struct {
	Token           string        `json:"token"`
	RequestCount    int           `json:"request_count"`
	HoldUntil       sql.NullInt64 `json:"hold_until"`
	LastRequestedAt sql.NullInt64 `json:"last_requested_at"`
}

// ResetPassword for
type ResetPassword struct {
	Otp             string        `json:"otp"`
	RequestCount    int           `json:"request_count"`
	FailedCount     int           `json:"failed_count"`
	HoldUntil       sql.NullInt64 `json:"hold_until"`
	IsValidated     bool          `json:"is_validated"`
	LastRequestedAt sql.NullInt64 `json:"last_requested_at"`
	TokenCreatedAt  sql.NullInt64 `json:"token_created_at"`
}

// Value for
func (vc VerificationCode) Value() (driver.Value, error) {
	return json.Marshal(vc)
}

// Scan for
func (vc *VerificationCode) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &vc)
}

// Value for
func (vc ResetPassword) Value() (driver.Value, error) {
	return json.Marshal(vc)
}

// Scan for
func (vc *ResetPassword) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &vc)
}

// Value for
func (vc EmailToken) Value() (driver.Value, error) {
	return json.Marshal(vc)
}

// Scan for
func (vc *EmailToken) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &vc)
}

// CreateUserVerificationDataParams for
type CreateUserVerificationDataParams struct {
	UserID           uuid.UUID        `json:"user_id"`
	VerificationCode VerificationCode `json:"verification_code"`
	EmailToken       EmailToken       `json:"email_token"`
	ResetPassword    ResetPassword    `json:"reset_password"`
	CreatedAt        sql.NullInt64    `json:"created_at"`
	CreatedBy        uuid.UUID        `json:"created_by"`
	UpdatedAt        sql.NullInt64    `json:"updated_at"`
	UpdatedBy        uuid.UUID        `json:"updated_by"`
	DeletedAt        sql.NullInt64    `json:"deleted_at"`
	DeletedBy        uuid.UUID        `json:"deleted_by"`
}

// UpdateVerificationCodeByUserIDParams is
type UpdateVerificationCodeByUserIDParams struct {
	UserID           uuid.UUID        `json:"user_id"`
	VerificationCode VerificationCode `json:"verification_code"`
	UpdatedAt        sql.NullInt64    `json:"updated_at"`
	UpdatedBy        uuid.UUID        `json:"updated_by"`
}

// UpdateEmailTokenByUserIDParams is
type UpdateEmailTokenByUserIDParams struct {
	UserID     uuid.UUID     `json:"user_id"`
	EmailToken EmailToken    `json:"email_token"`
	UpdatedAt  sql.NullInt64 `json:"updated_at"`
	UpdatedBy  uuid.UUID     `json:"updated_by"`
}

// UpdateResetPasswordByUserIDParams for
type UpdateResetPasswordByUserIDParams struct {
	UserID        uuid.UUID     `json:"user_id"`
	ResetPassword ResetPassword `json:"reset_password"`
	UpdatedAt     sql.NullInt64 `json:"updated_at"`
	UpdatedBy     uuid.UUID     `json:"updated_by"`
}
