package constant

import "fmt"

var (
	// ErrUserType is
	ErrUserType = fmt.Errorf("company account is not allowed")
	// ErrCompanyType is
	ErrCompanyType = fmt.Errorf("user account is not allowed")
	// ErrUserNotFound is
	ErrUserNotFound = fmt.Errorf("account not found")
	// ErrVerficationDataNotFound is
	ErrVerficationDataNotFound = fmt.Errorf("user verification data not found")
	// ErrWrongPassword is
	ErrWrongPassword = fmt.Errorf("invalid password")
	// ErrWrongOTPCode is
	ErrWrongOTPCode = fmt.Errorf("invalid otp code")
	// ErrEmailHasVerified is
	ErrEmailHasVerified = fmt.Errorf("your email has been verified")
	// ErrCannotSendEmail is
	ErrCannotSendEmail = fmt.Errorf("cannot send email, please try again later")
	// ErrOnHoldSendEmail is
	ErrOnHoldSendEmail = fmt.Errorf("cannot send email, your account is temporarily suspended")
	// ErrOnHoldOTPInput is
	ErrOnHoldOTPInput = fmt.Errorf("your account is temporarily suspended")
	// ErrCannotSendOTP is
	ErrCannotSendOTP = fmt.Errorf("cannot send otp code, please try again later")
	// ErrMsisdnHasVerified is
	ErrMsisdnHasVerified = fmt.Errorf("your phone number has been verified")
	// ErrExpiredOTP is
	ErrExpiredOTP = fmt.Errorf("otp code is expired")
	// ErrEmailNotVerified is
	ErrEmailNotVerified = fmt.Errorf("email address is not verified")
	// ErrMsisdnNotVerified is
	ErrMsisdnNotVerified = fmt.Errorf("phone number is not verified")
	// ErrCannotResetPassword is
	ErrCannotResetPassword = fmt.Errorf("cannot send reset password otp code, please try again later")
)
