package response

import "github.com/cpartogi/withdrawdeposit/entity"

// Register for
type Register struct {
	Email  string `json:"email,omitempty"`
	Msisdn string `json:"msisdn,omitempty"`
	Name   string `json:"name,omitempty"`
}

// Format is
func (ur *Register) Format(user entity.User) Register {
	ur.Email = user.Email
	ur.Msisdn = user.Msisdn.String
	ur.Name = user.Name

	return *ur
}
