package service

import (
	"context"
	"sync"
	"testing"

	"github.com/google/uuid"
)

func TestConcurrentDeposit(t *testing.T) {
	ctx := context.Background()

	// ⚠️ используем реальную БД (integration test)
	repo := setupTestRepo(t)
	service := NewWalletService(repo)

	walletID := uuid.New()

	// создаём кошелёк с балансом 0
	err := repo.CreateWallet(ctx, walletID)
	if err != nil {
		t.Fatalf("create wallet error: %v", err)
	}

	const workers = 100
	var wg sync.WaitGroup
	wg.Add(workers)

	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			err := service.Operate(ctx, walletID, "DEPOSIT", 1)
			if err != nil {
				t.Errorf("operate error: %v", err)
			}
		}()
	}

	wg.Wait()

	balance, err := service.GetBalance(ctx, walletID)
	if err != nil {
		t.Fatalf("get balance error: %v", err)
	}

	if balance != workers {
		t.Fatalf("expected balance %d, got %d", workers, balance)
	}
}
