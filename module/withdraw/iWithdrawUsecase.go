package withdraw

import (
	"context"

	"github.com/cpartogi/withdrawdeposit/entity"
	"github.com/cpartogi/withdrawdeposit/schema/response"
)

// Usecase is
type Usecase interface {
	RegisterUser(ctx context.Context, createUserParams entity.CreateUserParams) (user entity.User, err error)
	RegisterCompany(ctx context.Context, createUserParams entity.CreateUserParams) (user entity.User, err error)
	LoginUser(ctx context.Context, eUser entity.User) (token response.Token, err error)
	LoginOTP(ctx context.Context, eUser entity.User, otp string) (token response.Token, err error)
	LoginCompany(c context.Context, eUser entity.User) (token response.Token, err error)
	SendEmailVerification(ctx context.Context, email string) (user entity.User, err error)
	SendMisscallOTP(ctx context.Context, msisdn string) (misscallResponse response.MisscallOTP, err error)
	SendSMSOTP(ctx context.Context, msisdn, smsType, signature string) (smsResponse response.SMSViro, err error)
	SendResetPassword(ctx context.Context, email string) (user entity.User, err error)
	ValidateResetPassword(ctx context.Context, email, otp string) (user entity.User, err error)

	// legacy
	RegisterUserLegacy(ctx context.Context, createUserLegacyParams entity.CreateUserLegacyParams) (user entity.User, err error)
	RegisterCompanyLegacy(ctx context.Context, createUserLegacyParams entity.CreateUserLegacyParams) (user entity.User, err error)

	//deposit
	DepositBalance(ctx context.Context, seller_id string) (bal response.Balance, err error)
}
