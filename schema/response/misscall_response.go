package response

// MisscallOTP for
type MisscallOTP struct {
	RC      int    `json:"rc"`
	TrxID   string `json:"trxid"`
	Msisdn  string `json:"msisdn"`
	Token   string `json:"token"`
	Gateway int    `json:"gateway"`
}

// MisscallOTPNumber for
type MisscallOTPNumber struct {
	Number string `json:"number,omitempty"`
}

// Format is
func (ur *MisscallOTPNumber) Format(mOTP MisscallOTP) MisscallOTPNumber {
	rs := []rune(mOTP.Token)
	result := []rune{}
	for i := 0; i < len(rs)-4; i++ {
		result = append(result, rs[i])
	}
	ur.Number = string(result)

	return *ur
}
