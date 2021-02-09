package entity

type Seller struct {
	SellerName          string `validate:"required" json:"seller_name"`
	SellerEmail         string `validate:"required" json:"seller_email"`
	SellerBankCode      string `validate:"required" json:"seller_bank_code"`
	SellerAccountName   string `validate:"required" json:"seller_account_name"`
	SellerAccountNumber string `validate:"required" json:"seller_account_number"`
}

type SellerRow struct {
	Count int `json:"count"`
}
