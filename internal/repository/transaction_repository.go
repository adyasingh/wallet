package repository

import (
	"wallet-app/internal/errors"
	"wallet-app/internal/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	Save(transaction models.Transaction) (uint, error)
}

type TransactionRepositoryImpl struct {
	Db *gorm.DB
}

func NewTransactionRepositoryImpl(Db *gorm.DB) TransactionRepositoryImpl {
	return TransactionRepositoryImpl{Db: Db}
}

func (t TransactionRepositoryImpl) Save(transaction models.Transaction) (uint, error) {
	result := t.Db.Create(&transaction)
	if err := result.Error; err != nil {
		return 0, errors.HandleDbError(err)
	}
	return transaction.ID, nil
}
