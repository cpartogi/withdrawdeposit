package entity

type Balance struct {
	Amount   int    `validate:"required,number" json:"amount"`
	SellerId string `validate:"required" json:"seller_id"`
}

type BalanceRow struct {
	Count int `json:"count"`
}
