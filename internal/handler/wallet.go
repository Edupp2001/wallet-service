package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

type Service interface {
	Operate(ctx context.Context, id uuid.UUID, op string, amount int64) error
	GetBalance(ctx context.Context, id uuid.UUID) (int64, error)
}

type Handler struct {
	service Service
}

func New(s Service) *Handler {
	return &Handler{service: s}
}

type request struct {
	WalletID      uuid.UUID `json:"walletId"`
	OperationType string    `json:"operationType"`
	Amount        int64     `json:"amount"`
}

func (h *Handler) Wallet(w http.ResponseWriter, r *http.Request) {
	var req request
	json.NewDecoder(r.Body).Decode(&req)

	err := h.service.Operate(r.Context(), req.WalletID, req.OperationType, req.Amount)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) GetWallet(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/api/v1/wallets/"):]

	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "invalid wallet id", http.StatusBadRequest)
		return
	}

	balance, err := h.service.GetBalance(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"walletId": id,
		"balance":  balance,
	})
}
