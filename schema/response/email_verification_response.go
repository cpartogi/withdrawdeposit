package response

import "github.com/cpartogi/withdrawdeposit/entity"

// EmailVerification for
type EmailVerification struct {
	Email string `json:"email,omitempty"`
}

// Format is
func (ur *EmailVerification) Format(user entity.User) EmailVerification {
	ur.Email = user.Email

	return *ur
}
