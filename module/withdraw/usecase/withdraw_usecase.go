package usecase

import (
	"context"
	"time"

	"github.com/cpartogi/withdrawdeposit/entity"
	"github.com/cpartogi/withdrawdeposit/module/withdraw"
	"github.com/cpartogi/withdrawdeposit/schema/response"
)

// AuthUsecase will create a usecase with its required repo
type WithdrawUsecase struct {
	withdrawRepo   withdraw.Repository
	contextTimeout time.Duration
}

// NewAuthUsecase will create new an contactUsecase object representation of auth.Usecase
func NewWithdrawUsecase(ar withdraw.Repository, timeout time.Duration) withdraw.Usecase {
	return &WithdrawUsecase{
		withdrawRepo:   ar,
		contextTimeout: timeout,
	}
}

func (u *WithdrawUsecase) DepositBalance(ctx context.Context, seller_id string) (bal response.Balance, err error) {
	resp := response.Balance{}

	deposit, err := u.withdrawRepo.DepositBalance(ctx, seller_id)

	if err != nil {
		return resp, err
	}

	return deposit, err
}

func (u *WithdrawUsecase) DepositBalanceLog(ctx context.Context, seller_id string, date_from string, date_to string) (bal []response.BalanceLog, err error) {
	resp := []response.BalanceLog{}

	deposit, err := u.withdrawRepo.DepositBalanceLog(ctx, seller_id, date_from, date_to)

	if err != nil {
		return resp, err
	}

	return deposit, err
}

func (u *WithdrawUsecase) DepositRegister(ctx context.Context, depositRegister entity.Balance) (dep response.DepositRegistration, err error) {
	resp := response.DepositRegistration{
		SellerId: depositRegister.SellerId,
		Amount:   depositRegister.Amount,
	}

	req := entity.Balance{
		Amount:   depositRegister.Amount,
		SellerId: depositRegister.SellerId,
	}

	// cek if data exist
	sellerId := depositRegister.SellerId

	_, err = u.withdrawRepo.GetDepositBySellerid(ctx, sellerId)

	if err != nil {
		return resp, err
	}

	regdep, err := u.withdrawRepo.DepositRegister(ctx, req)

	if err != nil {
		return
	}

	return regdep, err
}

func (u *WithdrawUsecase) SellerRegister(ctx context.Context, sellerRegister entity.Seller) (sel response.SellerRegistration, err error) {
	resp := response.SellerRegistration{
		SellerName:          sellerRegister.SellerName,
		SellerEmail:         sellerRegister.SellerEmail,
		SellerBankCode:      sellerRegister.SellerBankCode,
		SellerAccountName:   sellerRegister.SellerAccountName,
		SellerAccountNumber: sellerRegister.SellerAccountName,
	}

	req := entity.Seller{
		SellerName:          sellerRegister.SellerName,
		SellerEmail:         sellerRegister.SellerEmail,
		SellerBankCode:      sellerRegister.SellerBankCode,
		SellerAccountName:   sellerRegister.SellerAccountName,
		SellerAccountNumber: sellerRegister.SellerAccountName,
	}

	//cek if seller exist
	sellerEmail := sellerRegister.SellerEmail

	_, err = u.withdrawRepo.GetSellerByEmail(ctx, sellerEmail)

	if err != nil {
		return resp, err
	}

	seller, err := u.withdrawRepo.SellerRegister(ctx, req)

	if err != nil {
		return
	}

	return seller, err

}

func (u *WithdrawUsecase) DisburseLog(ctx context.Context, transaction_id string, date_from string, date_to string) (dis []response.DisburseLog, err error) {
	resp := []response.DisburseLog{}

	disburse, err := u.withdrawRepo.DisburseLog(ctx, transaction_id, date_from, date_to)

	if err != nil {
		return resp, err
	}

	return disburse, err
}
