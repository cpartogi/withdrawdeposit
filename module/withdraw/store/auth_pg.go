package store

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/cpartogi/withdrawdeposit/entity"
	"github.com/cpartogi/withdrawdeposit/module/withdraw"
)

// SQLStore provides all functions to execute db queries and transactions.
type SQLStore struct {
	*Queries
	db *sql.DB
}

// NewStore creates a new store
func NewStore(db *sql.DB) withdraw.Repository {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

// execTX executes a function within a database transaction
func (s *SQLStore) execTX(ctx context.Context, fn func(*Queries) error) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

// RegisterTx is
func (s *SQLStore) RegisterTx(ctx context.Context, user entity.CreateUserParams, userVerification entity.CreateUserVerificationDataParams) (entity.User, error) {
	var result entity.User

	err := s.execTX(ctx, func(q *Queries) error {
		var err error

		result, err = q.CreateUser(ctx, user)
		if err != nil {
			return err
		}

		userVerification.UserID = result.ID

		_, err = q.CreateUserVerificationData(ctx, userVerification)
		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}

// LoginOTPTx is
func (s *SQLStore) LoginOTPTx(ctx context.Context,
	verificationCode entity.UpdateVerificationCodeByUserIDParams,
	userUpdateStatus entity.UpdateUserToVerifiedStatusParams) (entity.User, error) {

	var result entity.User

	err := s.execTX(ctx, func(q *Queries) error {
		var err error

		result, err = q.UpdateUserToVerifiedStatus(ctx, userUpdateStatus)
		if err != nil {
			return err
		}

		_, err = q.UpdateVerificationCodeByUserID(ctx, verificationCode)
		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}

// RegisterLegacyTx is
func (s *SQLStore) RegisterLegacyTx(ctx context.Context, user entity.CreateUserLegacyParams, userVerification entity.CreateUserVerificationDataParams) (entity.User, error) {
	var result entity.User

	err := s.execTX(ctx, func(q *Queries) error {
		var err error

		result, err = q.CreateUserLegacy(ctx, user)
		if err != nil {
			return err
		}

		userVerification.UserID = result.ID

		_, err = q.CreateUserVerificationData(ctx, userVerification)
		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}
