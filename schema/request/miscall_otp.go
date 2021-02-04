package request

// MisscallOTP for
type MisscallOTP struct {
	Msisdn string `validate:"required,number,startswith=62,min=10" json:"msisdn,omitempty"`
}

// CitcallMisscall is
type CitcallMisscall struct {
	Msisdn  string `json:"msisdn,omitempty"`
	Gateway string `json:"gateway,omitempty"`
}
