package utils

import (
	"log"
	"wallet-app/config"
	"wallet-app/internal/models"
	"wallet-app/internal/repository"

	"gorm.io/gorm"
)

var db *gorm.DB

func SetupTestDB() (*gorm.DB, *repository.TransactionRepositoryImpl, *repository.WalletRepositoryImpl, *repository.UserRepositoryImpl) {
	var err error
	databaseURL := "host=localhost port=5433 user=postgres password=postgres dbname=postgres sslmode=disable"
	db, err = config.DatabaseConnection(databaseURL)
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
