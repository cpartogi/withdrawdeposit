package withdraw

import (
	"context"

	"github.com/cpartogi/withdrawdeposit/entity"
	"github.com/cpartogi/withdrawdeposit/schema/response"
)

// Repository is
type Repository interface {
	//deposit
	DepositBalance(ctx context.Context, seller_id string) (bal response.Balance, err error)
	DepositBalanceLog(ctx context.Context, seller_id string, date_from string, date_to string) (bal []response.BalanceLog, err error)
	DepositRegister(ctx context.Context, arg entity.Balance) (dep response.DepositRegistration, err error)
	GetDepositBySellerid(ctx context.Context, seller_id string) (entity.BalanceRow, error)
	GetSellerByEmail(ctx context.Context, email string) (entity.SellerRow, error)
	SellerRegister(ctx context.Context, arg entity.Seller) (sel response.SellerRegistration, err error)
}
