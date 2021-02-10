package response

type DisburseLog struct {
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
