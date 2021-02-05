package request

type DepositRegistration struct {
	SellerId string `validate:"required" json:"seller_id,omitempty"`
	Amount   int    `validate:"required,number" json:"amount,omitempty"`
}
