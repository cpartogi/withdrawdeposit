package auth

import (
	"context"
	"database/sql"

	"github.com/cpartogi/withdrawdeposit/entity"
	"github.com/google/uuid"
)

// Repository is
type Repository interface {
	CreateUser(ctx context.Context, arg entity.CreateUserParams) (entity.User, error)
	CreateUserVerificationData(ctx context.Context, arg entity.CreateUserVerificationDataParams) (entity.UserVerification, error)
	GetUserByEmail(ctx context.Context, email string) (entity.User, error)
	GetUserByMsisdn(ctx context.Context, msisdn sql.NullString) (entity.User, error)
	GetVerificationCodeByUserID(ctx context.Context, userID uuid.UUID) (entity.VerificationCode, error)
	UpdateVerificationCodeByUserID(ctx context.Context, arg entity.UpdateVerificationCodeByUserIDParams) (entity.VerificationCode, error)
	GetEmailTokenByUserID(ctx context.Context, userID uuid.UUID) (entity.EmailToken, error)
	UpdateEmailTokenByUserID(ctx context.Context, arg entity.UpdateEmailTokenByUserIDParams) (entity.EmailToken, error)
	UpdateUserToVerifiedStatus(ctx context.Context, arg entity.UpdateUserToVerifiedStatusParams) (entity.User, error)
	GetResetPasswordByUserID(ctx context.Context, userID uuid.UUID) (entity.ResetPassword, error)
	UpdateResetPasswordByUserID(ctx context.Context, arg entity.UpdateResetPasswordByUserIDParams) (entity.ResetPassword, error)

	RegisterTx(ctx context.Context, user entity.CreateUserParams, userVerification entity.CreateUserVerificationDataParams) (entity.User, error)
	LoginOTPTx(ctx context.Context, verificationCode entity.UpdateVerificationCodeByUserIDParams, userMsisdn entity.UpdateUserToVerifiedStatusParams) (entity.User, error)

	// Legacy
	CreateUserLegacy(ctx context.Context, arg entity.CreateUserLegacyParams) (entity.User, error)
	RegisterLegacyTx(ctx context.Context, user entity.CreateUserLegacyParams, userVerification entity.CreateUserVerificationDataParams) (entity.User, error)
}
