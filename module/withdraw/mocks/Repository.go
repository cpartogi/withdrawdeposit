// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/cpartogi/withdrawdeposit/entity"
	mock "github.com/stretchr/testify/mock"

	sql "database/sql"

	response "github.com/cpartogi/withdrawdeposit/schema/response"
	uuid "github.com/google/uuid"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// CreateUser provides a mock function with given fields: ctx, arg
func (_m *Repository) CreateUser(ctx context.Context, arg entity.CreateUserParams) (entity.User, error) {
	ret := _m.Called(ctx, arg)

	var r0 entity.User
	if rf, ok := ret.Get(0).(func(context.Context, entity.CreateUserParams) entity.User); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(entity.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, entity.CreateUserParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateUserLegacy provides a mock function with given fields: ctx, arg
func (_m *Repository) CreateUserLegacy(ctx context.Context, arg entity.CreateUserLegacyParams) (entity.User, error) {
	ret := _m.Called(ctx, arg)

	var r0 entity.User
	if rf, ok := ret.Get(0).(func(context.Context, entity.CreateUserLegacyParams) entity.User); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(entity.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, entity.CreateUserLegacyParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateUserVerificationData provides a mock function with given fields: ctx, arg
func (_m *Repository) CreateUserVerificationData(ctx context.Context, arg entity.CreateUserVerificationDataParams) (entity.UserVerification, error) {
	ret := _m.Called(ctx, arg)

	var r0 entity.UserVerification
	if rf, ok := ret.Get(0).(func(context.Context, entity.CreateUserVerificationDataParams) entity.UserVerification); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(entity.UserVerification)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, entity.CreateUserVerificationDataParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEmailTokenByUserID provides a mock function with given fields: ctx, userID
func (_m *Repository) GetEmailTokenByUserID(ctx context.Context, userID uuid.UUID) (entity.EmailToken, error) {
	ret := _m.Called(ctx, userID)

	var r0 entity.EmailToken
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) entity.EmailToken); ok {
		r0 = rf(ctx, userID)
	} else {
		r0 = ret.Get(0).(entity.EmailToken)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetResetPasswordByUserID provides a mock function with given fields: ctx, userID
func (_m *Repository) GetResetPasswordByUserID(ctx context.Context, userID uuid.UUID) (entity.ResetPassword, error) {
	ret := _m.Called(ctx, userID)

	var r0 entity.ResetPassword
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) entity.ResetPassword); ok {
		r0 = rf(ctx, userID)
	} else {
		r0 = ret.Get(0).(entity.ResetPassword)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByEmail provides a mock function with given fields: ctx, email
func (_m *Repository) GetUserByEmail(ctx context.Context, email string) (entity.User, error) {
	ret := _m.Called(ctx, email)

	var r0 entity.User
	if rf, ok := ret.Get(0).(func(context.Context, string) entity.User); ok {
		r0 = rf(ctx, email)
	} else {
		r0 = ret.Get(0).(entity.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByMsisdn provides a mock function with given fields: ctx, msisdn
func (_m *Repository) GetUserByMsisdn(ctx context.Context, msisdn sql.NullString) (entity.User, error) {
	ret := _m.Called(ctx, msisdn)

	var r0 entity.User
	if rf, ok := ret.Get(0).(func(context.Context, sql.NullString) entity.User); ok {
		r0 = rf(ctx, msisdn)
	} else {
		r0 = ret.Get(0).(entity.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, sql.NullString) error); ok {
		r1 = rf(ctx, msisdn)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetVerificationCodeByUserID provides a mock function with given fields: ctx, userID
func (_m *Repository) GetVerificationCodeByUserID(ctx context.Context, userID uuid.UUID) (entity.VerificationCode, error) {
	ret := _m.Called(ctx, userID)

	var r0 entity.VerificationCode
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) entity.VerificationCode); ok {
		r0 = rf(ctx, userID)
	} else {
		r0 = ret.Get(0).(entity.VerificationCode)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LoginOTPTx provides a mock function with given fields: ctx, verificationCode, userMsisdn
func (_m *Repository) LoginOTPTx(ctx context.Context, verificationCode entity.UpdateVerificationCodeByUserIDParams, userMsisdn entity.UpdateUserToVerifiedStatusParams) (entity.User, error) {
	ret := _m.Called(ctx, verificationCode, userMsisdn)

	var r0 entity.User
	if rf, ok := ret.Get(0).(func(context.Context, entity.UpdateVerificationCodeByUserIDParams, entity.UpdateUserToVerifiedStatusParams) entity.User); ok {
		r0 = rf(ctx, verificationCode, userMsisdn)
	} else {
		r0 = ret.Get(0).(entity.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, entity.UpdateVerificationCodeByUserIDParams, entity.UpdateUserToVerifiedStatusParams) error); ok {
		r1 = rf(ctx, verificationCode, userMsisdn)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterLegacyTx provides a mock function with given fields: ctx, user, userVerification
func (_m *Repository) RegisterLegacyTx(ctx context.Context, user entity.CreateUserLegacyParams, userVerification entity.CreateUserVerificationDataParams) (entity.User, error) {
	ret := _m.Called(ctx, user, userVerification)

	var r0 entity.User
	if rf, ok := ret.Get(0).(func(context.Context, entity.CreateUserLegacyParams, entity.CreateUserVerificationDataParams) entity.User); ok {
		r0 = rf(ctx, user, userVerification)
	} else {
		r0 = ret.Get(0).(entity.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, entity.CreateUserLegacyParams, entity.CreateUserVerificationDataParams) error); ok {
		r1 = rf(ctx, user, userVerification)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterTx provides a mock function with given fields: ctx, user, userVerification
func (_m *Repository) RegisterTx(ctx context.Context, user entity.CreateUserParams, userVerification entity.CreateUserVerificationDataParams) (entity.User, error) {
	ret := _m.Called(ctx, user, userVerification)

	var r0 entity.User
	if rf, ok := ret.Get(0).(func(context.Context, entity.CreateUserParams, entity.CreateUserVerificationDataParams) entity.User); ok {
		r0 = rf(ctx, user, userVerification)
	} else {
		r0 = ret.Get(0).(entity.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, entity.CreateUserParams, entity.CreateUserVerificationDataParams) error); ok {
		r1 = rf(ctx, user, userVerification)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateEmailTokenByUserID provides a mock function with given fields: ctx, arg
func (_m *Repository) UpdateEmailTokenByUserID(ctx context.Context, arg entity.UpdateEmailTokenByUserIDParams) (entity.EmailToken, error) {
	ret := _m.Called(ctx, arg)

	var r0 entity.EmailToken
	if rf, ok := ret.Get(0).(func(context.Context, entity.UpdateEmailTokenByUserIDParams) entity.EmailToken); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(entity.EmailToken)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, entity.UpdateEmailTokenByUserIDParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateResetPasswordByUserID provides a mock function with given fields: ctx, arg
func (_m *Repository) UpdateResetPasswordByUserID(ctx context.Context, arg entity.UpdateResetPasswordByUserIDParams) (entity.ResetPassword, error) {
	ret := _m.Called(ctx, arg)

	var r0 entity.ResetPassword
	if rf, ok := ret.Get(0).(func(context.Context, entity.UpdateResetPasswordByUserIDParams) entity.ResetPassword); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(entity.ResetPassword)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, entity.UpdateResetPasswordByUserIDParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUserToVerifiedStatus provides a mock function with given fields: ctx, arg
func (_m *Repository) UpdateUserToVerifiedStatus(ctx context.Context, arg entity.UpdateUserToVerifiedStatusParams) (entity.User, error) {
	ret := _m.Called(ctx, arg)

	var r0 entity.User
	if rf, ok := ret.Get(0).(func(context.Context, entity.UpdateUserToVerifiedStatusParams) entity.User); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(entity.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, entity.UpdateUserToVerifiedStatusParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateVerificationCodeByUserID provides a mock function with given fields: ctx, arg
func (_m *Repository) UpdateVerificationCodeByUserID(ctx context.Context, arg entity.UpdateVerificationCodeByUserIDParams) (entity.VerificationCode, error) {
	ret := _m.Called(ctx, arg)

	var r0 entity.VerificationCode
	if rf, ok := ret.Get(0).(func(context.Context, entity.UpdateVerificationCodeByUserIDParams) entity.VerificationCode); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(entity.VerificationCode)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, entity.UpdateVerificationCodeByUserIDParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//deposit balance
func (_m *Repository) DepositBalance(ctx context.Context, seller_id string) (bal response.Balance, err error) {
	ret := _m.Called(ctx, seller_id)

	var r0 response.Balance
	if rf, ok := ret.Get(0).(func(context.Context, string) response.Balance); ok {
		r0 = rf(ctx, seller_id)
	} else {
		r0 = ret.Get(0).(response.Balance)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, seller_id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
