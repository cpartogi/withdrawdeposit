package response

// Login for
type Login struct {
	Token        string `json:"token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

// Format is
func (ur *Login) Format(token Token) Login {
	ur.Token = token.Token
	ur.RefreshToken = token.RefreshToken

	return *ur
}
