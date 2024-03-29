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

type SwaggerDepositBalanceLog struct {
	Base
	Data []DataBalanceLog `json:"data"`
}

type DataBalanceLog struct {
	LogDescription string `json:"log_description"`
	DepositBefore  int    `json:"deposit_before"`
	Amount         int    `json:"amount"`
	UpdatedDate    string `json:"updated_date"`
	UpdatedBy      string `json:"updated_by"`
}

type SwaggerDepositRegister struct {
	Base
	Data struct {
		Amount   int    `json:"amount"`
		SellerId string `json:"seller_id"`
	} `json:"data"`
}

type SwaggerSellerRegister struct {
	Base
	Data struct {
		SellerName          string `json:"seller_name"`
		SellerEmail         string `json:"seller_email"`
		SellerBankCode      string `json:"seller_bank_code"`
		SellerAccountName   string `json:"seller_account_name"`
		SellerAccountNumber string `json:"seller_account_number"`
	} `json:"data"`
}

type SwaggerDisburseLog struct {
	Base
	Data []DataDisburseLog `json:"data"`
}

type DataDisburseLog struct {
	LogId           string `json:"log_id"`
	TransactionId   int    `json:"transaction_id"`
	Amount          int    `json:"amount"`
	Fee             int    `json:"fee"`
	Remark          string `json:"remark"`
	Status          string `json:"status"`
	Receipt         string `json:"receipt"`
	BankCode        string `json:"bank_code"`
	AccountNumber   string `json:"account_number"`
	BeneficiaryName string `json:"beneficiary_name"`
	TimeServed      string `json:"time_served"`
	Timestamp       string `json:"timestamp"`
	CreatedDate     string `json:"created_date"`
}
