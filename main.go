package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"wallet-app/config"
	"wallet-app/internal/api/handlers"
	"wallet-app/internal/models"
	"wallet-app/internal/repository"
	"wallet-app/internal/router"
	"wallet-app/internal/service"
)

func main() {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	db, err := config.DatabaseConnection(databaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	err = db.AutoMigrate(&models.User{}, &models.Wallet{}, &models.Transaction{})

	if err != nil {
		log.Fatal("failed to migrate:", err)
	}

	txnRepo := repository.NewTransactionRepositoryImpl(db)
	walletRepo := repository.NewWalletRepositoryImpl(db)
	userRepo := repository.NewUserRepositoryImpl(db)

	walletSvc := service.NewWalletServiceImpl(txnRepo, walletRepo)
	walletController := handlers.NewWalletHandler(walletSvc)

	routes := router.NewRouter(walletController)

	// Seed Data
	user1, err := userRepo.Save(models.User{Name: "John Doe"})
	if err != nil {
		log.Fatal("failed to seed user:", err)
	}
	user2, err := userRepo.Save(models.User{Name: "Jane Smith"})
	if err != nil {
		log.Fatal("failed to seed user:", err)
	}
	walletRepo.Save(models.Wallet{UserID: user1, Balance: 100.0})
	walletRepo.Save(models.Wallet{UserID: user2, Balance: 200.0})

	server := &http.Server{
		Addr:           ":8080",
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	server.ListenAndServe()
}
