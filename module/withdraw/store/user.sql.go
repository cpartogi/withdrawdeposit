package store

import (
	"context"
	"database/sql"

	"github.com/cpartogi/withdrawdeposit/constant"
	"github.com/cpartogi/withdrawdeposit/entity"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
    name,
    email,
    msisdn,
    password,
    user_status_id,
    profile_image,
    gender,
    user_type_id,
    created_at,
    created_by,
    updated_at,
    updated_by
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
    $10,
    $11,
    $12
) RETURNING id, name, email, msisdn, password, user_status_id, profile_image, gender, user_type_id, is_msisdn_verified, is_email_verified, subscribe_id, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by
`

// CreateUser will
func (q *Queries) CreateUser(ctx context.Context, arg entity.CreateUserParams) (entity.User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.Name,
		arg.Email,
		arg.Msisdn,
		arg.Password,
		arg.UserStatusID,
		arg.ProfileImage,
		arg.Gender,
		arg.UserTypeID,
		arg.CreatedAt,
		arg.CreatedBy,
		arg.UpdatedAt,
		arg.UpdatedBy,
	)
	var i entity.User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Msisdn,
		&i.Password,
		&i.UserStatusID,
		&i.ProfileImage,
		&i.Gender,
		&i.UserTypeID,
		&i.IsMsisdnVerified,
		&i.IsEmailVerified,
		&i.SubscribeID,
		&i.CreatedAt,
		&i.CreatedBy,
		&i.UpdatedAt,
		&i.UpdatedBy,
		&i.DeletedAt,
		&i.DeletedBy,
	)
	return i, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, external_user_id, name, email, msisdn, password, user_status_id, profile_image, gender, user_type_id, is_msisdn_verified, is_email_verified, subscribe_id, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by from users
WHERE email = $1
`

// GetUserByEmail is
func (q *Queries) GetUserByEmail(ctx context.Context, email string) (entity.User, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i entity.User
	err := row.Scan(
		&i.ID,
		&i.ExternalUserID,
		&i.Name,
		&i.Email,
		&i.Msisdn,
		&i.Password,
		&i.UserStatusID,
		&i.ProfileImage,
		&i.Gender,
		&i.UserTypeID,
		&i.IsMsisdnVerified,
		&i.IsEmailVerified,
		&i.SubscribeID,
		&i.CreatedAt,
		&i.CreatedBy,
		&i.UpdatedAt,
		&i.UpdatedBy,
		&i.DeletedAt,
		&i.DeletedBy,
	)
	if err == sql.ErrNoRows {
		err = constant.ErrUserNotFound
	}

	return i, err
}

const getUserByMsisdn = `-- name: GetUserByMsisdn :one
SELECT id, external_user_id, name, email, msisdn, password, user_status_id, profile_image, gender, user_type_id, is_msisdn_verified, is_email_verified, subscribe_id, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by from users
WHERE msisdn = $1
`

// GetUserByMsisdn is
func (q *Queries) GetUserByMsisdn(ctx context.Context, msisdn sql.NullString) (entity.User, error) {
	row := q.db.QueryRowContext(ctx, getUserByMsisdn, msisdn)
	var i entity.User
	err := row.Scan(
		&i.ID,
		&i.ExternalUserID,
		&i.Name,
		&i.Email,
		&i.Msisdn,
		&i.Password,
		&i.UserStatusID,
		&i.ProfileImage,
		&i.Gender,
		&i.UserTypeID,
		&i.IsMsisdnVerified,
		&i.IsEmailVerified,
		&i.SubscribeID,
		&i.CreatedAt,
		&i.CreatedBy,
		&i.UpdatedAt,
		&i.UpdatedBy,
		&i.DeletedAt,
		&i.DeletedBy,
	)

	if err == sql.ErrNoRows {
		err = constant.ErrUserNotFound
	}

	return i, err
}

const updateUserToVerifiedStatus = `-- name: UpdateUserToVerifiedStatus :one
UPDATE users SET
is_msisdn_verified = $2,
user_status_id = $3
WHERE id = $1
RETURNING id, name, email, msisdn, password, user_status_id, profile_image, gender, user_type_id, is_msisdn_verified, is_email_verified, subscribe_id, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by
`

// UpdateUserToVerifiedStatus is
func (q *Queries) UpdateUserToVerifiedStatus(ctx context.Context, arg entity.UpdateUserToVerifiedStatusParams) (entity.User, error) {
	row := q.db.QueryRowContext(ctx, updateUserToVerifiedStatus, arg.ID, arg.IsMsisdnVerified, arg.UserStatusID)
	var i entity.User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Msisdn,
		&i.Password,
		&i.UserStatusID,
		&i.ProfileImage,
		&i.Gender,
		&i.UserTypeID,
		&i.IsMsisdnVerified,
		&i.IsEmailVerified,
		&i.SubscribeID,
		&i.CreatedAt,
		&i.CreatedBy,
		&i.UpdatedAt,
		&i.UpdatedBy,
		&i.DeletedAt,
		&i.DeletedBy,
	)
	if err == sql.ErrNoRows {
		err = constant.ErrUserNotFound
	}
	return i, err
}

const createUserLegacy = `-- name: CreateUserLegacy :one
INSERT INTO users (
    external_user_id,
    name,
    email,
    msisdn,
    password,
    user_status_id,
    profile_image,
    gender,
    user_type_id,
    created_at,
    created_by,
    updated_at,
    updated_by
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
    $10,
    $11,
    $12,
    $13
) RETURNING id, external_user_id, name, email, msisdn, password, user_status_id, profile_image, gender, user_type_id, is_msisdn_verified, is_email_verified, subscribe_id, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by
`

// CreateUserLegacy for
func (q *Queries) CreateUserLegacy(ctx context.Context, arg entity.CreateUserLegacyParams) (entity.User, error) {
	row := q.db.QueryRowContext(ctx, createUserLegacy,
		arg.ExternalUserID,
		arg.Name,
		arg.Email,
		arg.Msisdn,
		arg.Password,
		arg.UserStatusID,
		arg.ProfileImage,
		arg.Gender,
		arg.UserTypeID,
		arg.CreatedAt,
		arg.CreatedBy,
		arg.UpdatedAt,
		arg.UpdatedBy,
	)
	var i entity.User
	err := row.Scan(
		&i.ID,
		&i.ExternalUserID,
		&i.Name,
		&i.Email,
		&i.Msisdn,
		&i.Password,
		&i.UserStatusID,
		&i.ProfileImage,
		&i.Gender,
		&i.UserTypeID,
		&i.IsMsisdnVerified,
		&i.IsEmailVerified,
		&i.SubscribeID,
		&i.CreatedAt,
		&i.CreatedBy,
		&i.UpdatedAt,
		&i.UpdatedBy,
		&i.DeletedAt,
		&i.DeletedBy,
	)
	return i, err
}
