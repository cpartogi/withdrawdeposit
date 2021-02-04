package request

// ResetPassword for
type ResetPassword struct {
	Email string `validate:"required,email" json:"email,omitempty"`
}

// ValidateResetPassword for
type ValidateResetPassword struct {
	Email string `validate:"required,email" json:"email,omitempty"`
	OTP   string `validate:"required,len=5,number" json:"otp,omitempty"`
}
