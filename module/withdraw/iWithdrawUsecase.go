package withdraw

import (
	"context"

	"github.com/cpartogi/withdrawdeposit/entity"
	"github.com/cpartogi/withdrawdeposit/schema/response"
)

// Usecase is
type Usecase interface {

	//deposit
	DepositBalance(ctx context.Context, seller_id string) (bal response.Balance, err error)
	DepositBalanceLog(ctx context.Context, seller_id string, date_from string, date_to string) (bal []response.BalanceLog, err error)
	DepositRegister(ctx context.Context, depositRegister entity.Balance) (dep response.DepositRegistration, err error)
	SellerRegister(ctx context.Context, sellerRegister entity.Seller) (sel response.SellerRegistration, err error)
	DisburseLog(ctx context.Context, transaction_id string, date_from string, date_to string) (dis []response.DisburseLog, err error)
}
