package store

import (
	"context"
	"database/sql"
)

// DBTX will
type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

// New will
func New(db DBTX) *Queries {
	return &Queries{db: db}
}

// Queries will
type Queries struct {
	db DBTX
}

// WithTx will
func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db: tx,
	}
}
