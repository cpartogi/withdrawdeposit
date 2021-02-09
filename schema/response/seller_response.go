package response

type SellerRegistration struct {
	SellerName          string `json:"seller_name"`
	SellerEmail         string `json:"seller_email"`
	SellerBankCode      string `json:"seller_bank_code"`
	SellerAccountName   string `json:"seller_account_name"`
	SellerAccountNumber string `json:"seller_account_number"`
}
