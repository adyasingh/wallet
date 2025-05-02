package utils

import (
	"log"
	"os"
	"wallet-app/config"
	"wallet-app/internal/models"
	"wallet-app/internal/repository"

	"gorm.io/gorm"
)

func SetupTestDB() (*gorm.DB, *repository.TransactionRepositoryImpl, *repository.WalletRepositoryImpl, *repository.UserRepositoryImpl) {
	databaseURL := os.Getenv("TEST_DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("TEST_DATABASE_URL is not set")
	}
	db, err := config.DatabaseConnection(databaseURL)
	if err != nil {
		panic("failed to connect to the test database")
	}
	err = db.AutoMigrate(&models.User{}, &models.Wallet{}, &models.Transaction{})

	if err != nil {
		log.Fatal("failed to migrate:", err)
	}
	transactionRepo := &repository.TransactionRepositoryImpl{Db: db}
	walletRepo := &repository.WalletRepositoryImpl{Db: db}
	userRepo := &repository.UserRepositoryImpl{Db: db}
	return db, transactionRepo, walletRepo, userRepo
}
