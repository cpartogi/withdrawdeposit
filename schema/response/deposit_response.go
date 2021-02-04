package response

import (
	"github.com/google/uuid"
)

type Balance struct {
	Balance  int       `json:"balance" db:"amount"`
	SellerId uuid.UUID `json:"seller_id" db:"seller_id"`
}

type BalanceLog struct {
	LogDescription string `json:"log_description" db:"log_description"`
	DepositBefore  int    `json:"deposit_before" db:"deposit_before"`
	Amount         int    `json:"amount" db:"amount"`
	UpdatedDate    string `json:"updated_date" db:"updated_date"`
	UpdatedBy      string `json:"updated_by" db:"updated_by"`
}
