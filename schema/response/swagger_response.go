package response

// SwaggerUserRegister for
type SwaggerUserRegister struct {
	Base
	Data struct {
		Email  string `json:"email,omitempty"`
		Msisdn string `json:"msisdn,omitempty"`
		Name   string `json:"name,omitempty"`
	} `json:"data"`
}

// SwaggerCompanyRegister for
type SwaggerCompanyRegister struct {
	Base
	Data struct {
		Email string `json:"email,omitempty"`
		Name  string `json:"name,omitempty"`
	} `json:"data"`
}

// SwaggerUserLogin for
type SwaggerUserLogin struct {
	Base
	Data struct {
		Token        string `json:"token,omitempty"`
		RefreshToken string `json:"refresh_token,omitempty"`
	} `json:"data"`
}

// SwaggerOTPLogin for
type SwaggerOTPLogin struct {
	Base
	Data struct {
		Token        string `json:"token,omitempty"`
		RefreshToken string `json:"refresh_token,omitempty"`
	} `json:"data"`
}

// SwaggerCompanyLogin for
type SwaggerCompanyLogin struct {
	Base
	Data struct {
		Token        string `json:"token,omitempty"`
		RefreshToken string `json:"refresh_token,omitempty"`
	} `json:"data"`
}

// SwaggerEmailVerification for
type SwaggerEmailVerification struct {
	Base
	Data struct {
		Email string `json:"email,omitempty"`
	} `json:"data"`
}

// SwaggerMisscallOTP for
type SwaggerMisscallOTP struct {
	Base
	Data struct {
		Number string `json:"number,omitempty"`
	} `json:"data"`
}

// SwaggerSMSOTP for
type SwaggerSMSOTP struct {
	Base
	Data struct {
		Msisdn string `json:"msisdn,omitempty"`
		Type   string `json:"type,omitempty"`
	} `json:"data"`
}

// SwaggerResetPassword for
type SwaggerResetPassword struct {
	Base
	Data struct {
		Email string `json:"email,omitempty"`
	} `json:"data"`
}

type SwaggerDepositBalance struct {
	Base
	Data struct {
		Balance  int    `json:"balance,omitempty"`
		SellerId string `json:"seller_id,omitempty"`
	} `json:"data"`
}
