package response

type Balance struct {
	Balance  int    `json:"balance"`
	SellerId string `json:"seller_id"`
}

type BalanceLog struct {
	LogDescription string `json:"log_description"`
	DepositBefore  int    `json:"deposit_before"`
	Amount         int    `json:"amount"`
	UpdatedDate    string `json:"updated_date"`
	UpdatedBy      string `json:"updated_by"`
}

type DepositRegistration struct {
	SellerId string `json:"seller_id,omitempty"`
	Amount   int    `json:"amount,omitempty"`
}
