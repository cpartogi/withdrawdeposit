package store

import (
	"context"
	"database/sql"

	"github.com/cpartogi/withdrawdeposit/constant"
	"github.com/cpartogi/withdrawdeposit/schema/response"
)

const getDepositBalance = `-- name: DepositBalance :one
SELECT amount, seller_id FROM tb_seller_deposit
WHERE seller_id = ?
`

// Deposit balance is
func (q *Queries) DepositBalance(ctx context.Context, seller_id string) (response.Balance, error) {
	row := q.db.QueryRowContext(ctx, getDepositBalance, seller_id)
	var i response.Balance
	err := row.Scan(
		&i.Balance,
		&i.SellerId,
	)
	if err == sql.ErrNoRows {
		err = constant.ErrUserNotFound
	}

	return i, err
}
