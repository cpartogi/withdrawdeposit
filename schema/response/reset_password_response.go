package response

import "github.com/cpartogi/withdrawdeposit/entity"

// ResetPassword for
type ResetPassword struct {
	Email string `json:"email,omitempty"`
}

// Format is
func (rp *ResetPassword) Format(user *entity.User) *ResetPassword {
	rp.Email = user.Email

	return rp
}

// ValidateResetPassword for
type ValidateResetPassword struct {
	Email string `json:"email,omitempty"`
	OTP   string `json:"otp,omitempty"`
}

// Format is
func (rp *ValidateResetPassword) Format(user *entity.User, otp string) *ValidateResetPassword {
	rp.Email = user.Email
	rp.OTP = otp

	return rp
}
