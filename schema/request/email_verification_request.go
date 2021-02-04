package request

// EmailVerification for
type EmailVerification struct {
	Email string `validate:"required,email" json:"email,omitempty"`
}
