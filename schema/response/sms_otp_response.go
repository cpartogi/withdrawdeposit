package response

// SMSOTP for
type SMSOTP struct {
	Msisdn string `json:"msisdn,omitempty"`
	Type   string `json:"type,omitempty"`
}

// Format is
func (so *SMSOTP) Format(smsViro SMSViro, smsType string) SMSOTP {
	so.Msisdn = smsViro.Messages[0].To
	so.Type = smsType

	return *so
}

// SMSViro for
type SMSViro struct {
	Messages []struct {
		To     string `json:"to,omitempty"`
		Status struct {
			GroupID     int    `json:"groupId"`
			GroupName   string `json:"groupName"`
			ID          int    `json:"id"`
			Name        string `json:"name"`
			Description string `json:"description"`
		} `json:"status,omitempty"`
		MessageID string `json:"messageId,omitempty"`
		SMSCount  int    `json:"smsCount,omitempty"`
	} `json:"messages,omitempty"`
}
