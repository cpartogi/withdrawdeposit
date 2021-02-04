// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/cpartogi/withdrawdeposit/entity"
	mock "github.com/stretchr/testify/mock"

	response "github.com/cpartogi/withdrawdeposit/schema/response"
)

// Usecase is an autogenerated mock type for the Usecase type
type Usecase struct {
	mock.Mock
}

// LoginCompany provides a mock function with given fields: c, eUser
func (_m *Usecase) LoginCompany(c context.Context, eUser entity.User) (response.Token, error) {
	ret := _m.Called(c, eUser)

	var r0 response.Token
	if rf, ok := ret.Get(0).(func(context.Context, entity.User) response.Token); ok {
		r0 = rf(c, eUser)
	} else {
		r0 = ret.Get(0).(response.Token)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, entity.User) error); ok {
		r1 = rf(c, eUser)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LoginOTP provides a mock function with given fields: ctx, eUser, otp
func (_m *Usecase) LoginOTP(ctx context.Context, eUser entity.User, otp string) (response.Token, error) {
	ret := _m.Called(ctx, eUser, otp)

	var r0 response.Token
	if rf, ok := ret.Get(0).(func(context.Context, entity.User, string) response.Token); ok {
		r0 = rf(ctx, eUser, otp)
	} else {
		r0 = ret.Get(0).(response.Token)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, entity.User, string) error); ok {
		r1 = rf(ctx, eUser, otp)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LoginUser provides a mock function with given fields: ctx, eUser
func (_m *Usecase) LoginUser(ctx context.Context, eUser entity.User) (response.Token, error) {
	ret := _m.Called(ctx, eUser)

	var r0 response.Token
	if rf, ok := ret.Get(0).(func(context.Context, entity.User) response.Token); ok {
		r0 = rf(ctx, eUser)
	} else {
		r0 = ret.Get(0).(response.Token)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, entity.User) error); ok {
		r1 = rf(ctx, eUser)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterCompany provides a mock function with given fields: ctx, createUserParams
func (_m *Usecase) RegisterCompany(ctx context.Context, createUserParams entity.CreateUserParams) (entity.User, error) {
	ret := _m.Called(ctx, createUserParams)

	var r0 entity.User
	if rf, ok := ret.Get(0).(func(context.Context, entity.CreateUserParams) entity.User); ok {
		r0 = rf(ctx, createUserParams)
	} else {
		r0 = ret.Get(0).(entity.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, entity.CreateUserParams) error); ok {
		r1 = rf(ctx, createUserParams)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterCompanyLegacy provides a mock function with given fields: ctx, createUserLegacyParams
func (_m *Usecase) RegisterCompanyLegacy(ctx context.Context, createUserLegacyParams entity.CreateUserLegacyParams) (entity.User, error) {
	ret := _m.Called(ctx, createUserLegacyParams)

	var r0 entity.User
	if rf, ok := ret.Get(0).(func(context.Context, entity.CreateUserLegacyParams) entity.User); ok {
		r0 = rf(ctx, createUserLegacyParams)
	} else {
		r0 = ret.Get(0).(entity.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, entity.CreateUserLegacyParams) error); ok {
		r1 = rf(ctx, createUserLegacyParams)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterUser provides a mock function with given fields: ctx, createUserParams
func (_m *Usecase) RegisterUser(ctx context.Context, createUserParams entity.CreateUserParams) (entity.User, error) {
	ret := _m.Called(ctx, createUserParams)

	var r0 entity.User
	if rf, ok := ret.Get(0).(func(context.Context, entity.CreateUserParams) entity.User); ok {
		r0 = rf(ctx, createUserParams)
	} else {
		r0 = ret.Get(0).(entity.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, entity.CreateUserParams) error); ok {
		r1 = rf(ctx, createUserParams)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterUserLegacy provides a mock function with given fields: ctx, createUserLegacyParams
func (_m *Usecase) RegisterUserLegacy(ctx context.Context, createUserLegacyParams entity.CreateUserLegacyParams) (entity.User, error) {
	ret := _m.Called(ctx, createUserLegacyParams)

	var r0 entity.User
	if rf, ok := ret.Get(0).(func(context.Context, entity.CreateUserLegacyParams) entity.User); ok {
		r0 = rf(ctx, createUserLegacyParams)
	} else {
		r0 = ret.Get(0).(entity.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, entity.CreateUserLegacyParams) error); ok {
		r1 = rf(ctx, createUserLegacyParams)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendEmailVerification provides a mock function with given fields: ctx, email
func (_m *Usecase) SendEmailVerification(ctx context.Context, email string) (entity.User, error) {
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

// SendMisscallOTP provides a mock function with given fields: ctx, msisdn
func (_m *Usecase) SendMisscallOTP(ctx context.Context, msisdn string) (response.MisscallOTP, error) {
	ret := _m.Called(ctx, msisdn)

	var r0 response.MisscallOTP
	if rf, ok := ret.Get(0).(func(context.Context, string) response.MisscallOTP); ok {
		r0 = rf(ctx, msisdn)
	} else {
		r0 = ret.Get(0).(response.MisscallOTP)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, msisdn)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendResetPassword provides a mock function with given fields: ctx, email
func (_m *Usecase) SendResetPassword(ctx context.Context, email string) (entity.User, error) {
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

// SendSMSOTP provides a mock function with given fields: ctx, msisdn, smsType, signature
func (_m *Usecase) SendSMSOTP(ctx context.Context, msisdn string, smsType string, signature string) (response.SMSViro, error) {
	ret := _m.Called(ctx, msisdn, smsType, signature)

	var r0 response.SMSViro
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) response.SMSViro); ok {
		r0 = rf(ctx, msisdn, smsType, signature)
	} else {
		r0 = ret.Get(0).(response.SMSViro)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string) error); ok {
		r1 = rf(ctx, msisdn, smsType, signature)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValidateResetPassword provides a mock function with given fields: ctx, email, otp
func (_m *Usecase) ValidateResetPassword(ctx context.Context, email string, otp string) (entity.User, error) {
	ret := _m.Called(ctx, email, otp)

	var r0 entity.User
	if rf, ok := ret.Get(0).(func(context.Context, string, string) entity.User); ok {
		r0 = rf(ctx, email, otp)
	} else {
		r0 = ret.Get(0).(entity.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, email, otp)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}