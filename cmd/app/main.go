package main

import (
	"log"
	"net/http"

	"wallet-service/internal/config"
	"wallet-service/internal/db"
	"wallet-service/internal/handler"
	"wallet-service/internal/repository"
	"wallet-service/internal/service"
)

func main() {
	// 1. Загружаем конфиг
	cfg := config.Load()

	// 2. Подключаемся к БД
	database, err := db.NewPostgres()
	if err != nil {
		log.Fatal(err)
	}

	// 3. Собираем зависимости
	walletRepo := repository.NewWalletRepo(database)
	walletService := service.NewWalletService(walletRepo)
	walletHandler := handler.New(walletService)

	// 4. Роутинг
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/wallet", walletHandler.Wallet)
	mux.HandleFunc("/api/v1/wallets/", walletHandler.GetWallet)

	// 5. Запуск сервера
	log.Println("Server started on :" + cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, mux))

}
