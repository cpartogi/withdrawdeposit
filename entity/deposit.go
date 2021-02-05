package entity

type Balance struct {
	Amount   int    `json:"amount"`
	SellerId string `json:"seller_id"`
}

type BalanceRow struct {
	Count int `json:"count"`
}
