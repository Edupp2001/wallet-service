package service

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	UpdateBalance(ctx context.Context, id uuid.UUID, delta int64) error
	GetBalance(ctx context.Context, id uuid.UUID) (int64, error)
	CreateWallet(ctx context.Context, id uuid.UUID) error
}

type WalletService struct {
	repo Repository
}

func NewWalletService(r Repository) *WalletService {
	return &WalletService{repo: r}
}

func (s *WalletService) Operate(ctx context.Context, id uuid.UUID, op string, amount int64) error {
	if op == "WITHDRAW" {
		amount = -amount
	}
	return s.repo.UpdateBalance(ctx, id, amount)
}

func (s *WalletService) GetBalance(ctx context.Context, id uuid.UUID) (int64, error) {
	return s.repo.GetBalance(ctx, id)
}
