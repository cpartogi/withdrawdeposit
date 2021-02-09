package store

import (
	"context"

	"github.com/cpartogi/withdrawdeposit/constant"
	"github.com/cpartogi/withdrawdeposit/entity"
	"github.com/cpartogi/withdrawdeposit/schema/response"
)

const getSellerByEmail = `-- name: SellerByEmail :one
SELECT count(seller_id) as row FROM tb_seller
WHERE seller_email = ?
`

func (q *Queries) GetSellerByEmail(ctx context.Context, seller_email string) (entity.SellerRow, error) {
	row := q.db.QueryRowContext(ctx, getSellerByEmail, seller_email)
	var i entity.SellerRow

	err := row.Scan(
		&i.Count,
	)

	totalrow := i.Count

	if totalrow > 0 {
		err = constant.ErrConflict
	}

	return i, err

}

const createSeller = `-- name: CreateSeller :one
INSERT INTO tb_seller (
    seller_id,
    seller_name,
    seller_email,
	seller_bank_code,
	seller_account_name,
	seller_account_number,
    created_by,
    updated_by
) VALUES (
    uuid(),
    ?,
    ?,
	?,
	?,
	?,
    'admin',
    'admin'
)
`

func (q *Queries) SellerRegister(ctx context.Context, arg entity.Seller) (sel response.SellerRegistration, err error) {

	result, err := q.db.ExecContext(ctx, createSeller,
		arg.SellerName,
		arg.SellerEmail,
		arg.SellerBankCode,
		arg.SellerAccountName,
		arg.SellerAccountNumber,
	)

	if err != nil {
		return
	}

	rows, err := result.RowsAffected()

	if rows != 1 {
		return
	}

	i := response.SellerRegistration{
		SellerName:          arg.SellerName,
		SellerEmail:         arg.SellerEmail,
		SellerBankCode:      arg.SellerBankCode,
		SellerAccountName:   arg.SellerBankCode,
		SellerAccountNumber: arg.SellerAccountNumber,
	}

	return i, err
}
