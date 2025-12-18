package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
)

type WalletRepo struct {
	db *sql.DB
}

func NewWalletRepo(db *sql.DB) *WalletRepo {
	return &WalletRepo{db: db}
}

func (r *WalletRepo) UpdateBalance(ctx context.Context, id uuid.UUID, delta int64) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var balance int64
	err = tx.QueryRowContext(ctx,
		"SELECT balance FROM wallets WHERE id=$1 FOR UPDATE",
		id,
	).Scan(&balance)
	if err != nil {
		return err
	}

	if balance+delta < 0 {
		return errors.New("insufficient funds")
	}

	_, err = tx.ExecContext(ctx,
		"UPDATE wallets SET balance = balance + $1 WHERE id=$2",
		delta, id,
	)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r *WalletRepo) GetBalance(ctx context.Context, id uuid.UUID) (int64, error) {
	var balance int64
	err := r.db.QueryRowContext(
		ctx,
		"SELECT balance FROM wallets WHERE id = $1",
		id,
	).Scan(&balance)

	if err == sql.ErrNoRows {
		return 0, errors.New("wallet not found")
	}
	return balance, err
}
