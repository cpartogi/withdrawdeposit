package store

import (
	"context"

	"github.com/cpartogi/withdrawdeposit/constant"

	"github.com/cpartogi/withdrawdeposit/schema/response"
)

const getDisburseLog = `-- name: DisburseLog
SELECT log_id, transaction_id, amount, fee, remark, status, receipt, bank_code, account_number, beneficiary_name, time_served, timestamp, created_date FROM tb_transaction_log
WHERE  transaction_id = ? AND created_date > ? AND created_date < ?  ORDER BY created_date DESC
`

// Deposit balancelog is
func (q *Queries) DisburseLog(ctx context.Context, transaction_id string, date_from string, date_to string) ([]response.DisburseLog, error) {
	rows, err := q.db.QueryContext(ctx, getDisburseLog, transaction_id, date_from, date_to)

	defer rows.Close()

	var y []response.DisburseLog
	var i response.DisburseLog

	c := 0
	for rows.Next() {
		_ = rows.Scan(
			&i.LogId,
			&i.TransactionId,
			&i.Amount,
			&i.Fee,
			&i.Remark,
			&i.Status,
			&i.Receipt,
			&i.BankCode,
			&i.AccountNumber,
			&i.BeneficiaryName,
			&i.TimeServed,
			&i.Timestamp,
			&i.CreatedDate,
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
