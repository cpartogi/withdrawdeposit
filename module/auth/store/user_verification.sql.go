package store

import (
	"context"
	"database/sql"

	"github.com/cpartogi/withdrawdeposit/constant"
	"github.com/cpartogi/withdrawdeposit/entity"
	"github.com/google/uuid"
)

const createUserVerificationData = `-- name: CreateUserVerificationData :one
INSERT INTO user_verifications (
   user_id,
   verification_code,
   email_token,
   reset_password,
   created_at,
   created_by,
   updated_at,
   updated_by,
   deleted_at,
   deleted_by
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8,
    $9,
    $10
)
RETURNING id, user_id, verification_code, email_token, reset_password, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by
`

// CreateUserVerificationData will
func (q *Queries) CreateUserVerificationData(ctx context.Context, arg entity.CreateUserVerificationDataParams) (entity.UserVerification, error) {
	row := q.db.QueryRowContext(ctx, createUserVerificationData,
		arg.UserID,
		arg.VerificationCode,
		arg.EmailToken,
		arg.ResetPassword,
		arg.CreatedAt,
		arg.CreatedBy,
		arg.UpdatedAt,
		arg.UpdatedBy,
		arg.DeletedAt,
		arg.DeletedBy,
	)
	var i entity.UserVerification
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.VerificationCode,
		&i.EmailToken,
		&i.ResetPassword,
		&i.CreatedAt,
		&i.CreatedBy,
		&i.UpdatedAt,
		&i.UpdatedBy,
		&i.DeletedAt,
		&i.DeletedBy,
	)
	return i, err
}

const getVerificationCodeByUserID = `-- name: GetVerificationCodeByUserID :one
SELECT verification_code from user_verifications
WHERE user_id = $1
`

// GetVerificationCodeByUserID returns
func (q *Queries) GetVerificationCodeByUserID(ctx context.Context, userID uuid.UUID) (entity.VerificationCode, error) {
	row := q.db.QueryRowContext(ctx, getVerificationCodeByUserID, userID)
	var verificationCode entity.VerificationCode
	err := row.Scan(&verificationCode)
	if err == sql.ErrNoRows {
		err = constant.ErrVerficationDataNotFound
	}
	return verificationCode, err
}

const updateVerificationCodeByUserID = `-- name: UpdateVerificationCodeByUserID :one
UPDATE user_verifications
SET verification_code = $2, 
updated_at = $3,
updated_by = $4
WHERE user_id = $1
RETURNING verification_code
`

// UpdateVerificationCodeByUserID will
func (q *Queries) UpdateVerificationCodeByUserID(ctx context.Context, arg entity.UpdateVerificationCodeByUserIDParams) (entity.VerificationCode, error) {
	row := q.db.QueryRowContext(ctx, updateVerificationCodeByUserID,
		arg.UserID,
		arg.VerificationCode,
		arg.UpdatedAt,
		arg.UpdatedBy,
	)
	var verificationCode entity.VerificationCode
	err := row.Scan(&verificationCode)
	if err == sql.ErrNoRows {
		err = constant.ErrVerficationDataNotFound
	}
	return verificationCode, err
}

const updateEmailTokenByUserID = `-- name: UpdateEmailTokenByUserID :one
UPDATE user_verifications SET email_token = $2,
updated_at = $3,
updated_by = $4
WHERE user_id = $1
RETURNING email_token
`

// UpdateEmailTokenByUserID is
func (q *Queries) UpdateEmailTokenByUserID(ctx context.Context, arg entity.UpdateEmailTokenByUserIDParams) (entity.EmailToken, error) {
	row := q.db.QueryRowContext(ctx, updateEmailTokenByUserID,
		arg.UserID,
		arg.EmailToken,
		arg.UpdatedAt,
		arg.UpdatedBy,
	)
	var emailToken entity.EmailToken
	err := row.Scan(&emailToken)
	if err == sql.ErrNoRows {
		err = constant.ErrVerficationDataNotFound
	}

	return emailToken, err
}

const getEmailTokenByUserID = `-- name: GetEmailTokenByUserID :one
SELECT email_token from user_verifications
WHERE user_id = $1
`

//GetEmailTokenByUserID is
func (q *Queries) GetEmailTokenByUserID(ctx context.Context, userID uuid.UUID) (entity.EmailToken, error) {
	row := q.db.QueryRowContext(ctx, getEmailTokenByUserID, userID)
	var emailToken entity.EmailToken
	err := row.Scan(&emailToken)
	if err == sql.ErrNoRows {
		err = constant.ErrVerficationDataNotFound
	}
	return emailToken, err
}

const getResetPasswordByUserID = `-- name: GetResetPasswordByUserID :one
SELECT reset_password from user_verifications
WHERE user_id = $1
`

// GetResetPasswordByUserID for
func (q *Queries) GetResetPasswordByUserID(ctx context.Context, userID uuid.UUID) (entity.ResetPassword, error) {
	row := q.db.QueryRowContext(ctx, getResetPasswordByUserID, userID)
	var resetPassword entity.ResetPassword
	err := row.Scan(&resetPassword)
	if err == sql.ErrNoRows {
		err = constant.ErrVerficationDataNotFound
	}
	return resetPassword, err
}

const updateResetPasswordByUserID = `-- name: UpdateResetPasswordByUserID :one
UPDATE user_verifications
SET reset_password = $2, 
updated_at = $3,
updated_by = $4
WHERE user_id = $1
RETURNING reset_password
`

// UpdateResetPasswordByUserID for
func (q *Queries) UpdateResetPasswordByUserID(ctx context.Context, arg entity.UpdateResetPasswordByUserIDParams) (entity.ResetPassword, error) {
	row := q.db.QueryRowContext(ctx, updateResetPasswordByUserID,
		arg.UserID,
		arg.ResetPassword,
		arg.UpdatedAt,
		arg.UpdatedBy,
	)
	var resetPassword entity.ResetPassword
	err := row.Scan(&resetPassword)
	if err == sql.ErrNoRows {
		err = constant.ErrVerficationDataNotFound
	}

	return resetPassword, err
}
