package request

// SMSOTP for
type SMSOTP struct {
	Msisdn    string `validate:"required,number,startswith=62,min=10" json:"msisdn,omitempty"`
	Type      string `validate:"required,oneof=login register" json:"type,omitempty"`
	Signature string `json:"signature,omitempty"`
}

// SmsViro is
type SmsViro struct {
	From string `json:"from"`
	To   string `json:"to"`
	Text string `json:"text"`
}
