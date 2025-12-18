package service

import (
	"testing"

	"wallet-service/internal/db"
	"wallet-service/internal/repository"
)

func setupTestRepo(t *testing.T) *repository.WalletRepo {
	t.Helper()

	database, err := db.NewPostgres()
	if err != nil {
		t.Fatal(err)
	}

	return repository.NewWalletRepo(database)
}
