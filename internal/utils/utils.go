package utils

import (
	"log"
	"os"
	"sync"
	"wallet-app/internal/models"
	"wallet-app/internal/repository"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

// GetDB initializes and returns the singleton instance of the database connection
func SetupTestDB() (*gorm.DB, *repository.TransactionRepositoryImpl, *repository.WalletRepositoryImpl, *repository.UserRepositoryImpl) {
	databaseURL := os.Getenv("TEST_DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("TEST_DATABASE_URL is not set")
	}
	once.Do(func() {
		var err error
		db, err = gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
		if err != nil {
			log.Fatal("Failed to connect to the database:", err)
		}

		err = db.AutoMigrate(&models.User{}, &models.Wallet{}, &models.Transaction{})
		if err != nil {
			log.Fatal("Failed to migrate:", err)
		}
	})
	transactionRepo := &repository.TransactionRepositoryImpl{Db: db}
	walletRepo := &repository.WalletRepositoryImpl{Db: db}
	userRepo := &repository.UserRepositoryImpl{Db: db}
	return db, transactionRepo, walletRepo, userRepo
}

// CloseDB closes the database connection
func CloseDB() {
	if db != nil {
		sqlDB, err := db.DB()
		if err != nil {
			log.Println("Failed to get DB instance:", err)
			return
		}
		sqlDB.Close()
	}
}
