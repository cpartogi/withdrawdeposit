package store

import (
	"context"
	"database/sql"

	"github.com/cpartogi/withdrawdeposit/constant"
	"github.com/cpartogi/withdrawdeposit/entity"
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

const getDepositBalanceLog = `-- name: DepositBalanceLog
SELECT a.log_description, a.deposit_before, a.amount, a.updated_date, a.updated_by FROM tb_seller_deposit_log a, tb_seller_deposit b
WHERE a.deposit_id=b.deposit_id AND b.seller_id = ? AND a.updated_date > ? AND a.updated_date < ?  ORDER BY a.updated_date DESC
`

// Deposit balancelog is
func (q *Queries) DepositBalanceLog(ctx context.Context, seller_id string, date_from string, date_to string) ([]response.BalanceLog, error) {
	rows, err := q.db.QueryContext(ctx, getDepositBalanceLog, seller_id, date_from, date_to)

	defer rows.Close()

	var y []response.BalanceLog
	var i response.BalanceLog

	c := 0
	for rows.Next() {
		_ = rows.Scan(
			&i.LogDescription,
			&i.DepositBefore,
			&i.Amount,
			&i.UpdatedDate,
			&i.UpdatedBy,
		)
		y = append(y, i)
		c++
	}

	//return not found
	if c == 0 {
		err = constant.ErrUserNotFound
	}

	return y, err
}

const getDepositBySellerId = `-- name: DepositBySellerId :one
SELECT count(seller_id) as row FROM tb_seller_deposit
WHERE seller_id = ?
`

// Deposit balance is
func (q *Queries) GetDepositBySellerid(ctx context.Context, seller_id string) (entity.BalanceRow, error) {
	row := q.db.QueryRowContext(ctx, getDepositBySellerId, seller_id)
	var i entity.BalanceRow
	err := row.Scan(
		&i.Count,
	)

	totalrow := i.Count

	if totalrow > 0 {
		err = constant.ErrConflict
	}

	return i, err
}

const createDeposit = `-- name: CreateDeposit :one
INSERT INTO tb_seller_deposit (
    deposit_id,
    seller_id,
    amount,
    created_by,
    updated_by
) VALUES (
    uuid(),
    ?,
    ?,
    'admin',
    'admin'
)
`

func (q *Queries) DepositRegister(ctx context.Context, arg entity.Balance) (dep response.DepositRegistration, err error) {

	result, err := q.db.ExecContext(ctx, createDeposit,
		arg.SellerId,
		arg.Amount,
	)

	if err != nil {
		return
	}

	rows, err := result.RowsAffected()

	if err != nil {
		return
	}

	if rows != 1 {
		return
	}

	i := response.DepositRegistration{
		SellerId: arg.SellerId,
		Amount:   arg.Amount,
	}

	return i, err
}
