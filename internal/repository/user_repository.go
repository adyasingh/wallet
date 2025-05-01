package repository

import (
	"wallet-app/internal/errors"
	"wallet-app/internal/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	Save(transaction models.User) (uint, error)
}

type UserRepositoryImpl struct {
	Db *gorm.DB
}

func NewUserRepositoryImpl(Db *gorm.DB) UserRepositoryImpl {
	return UserRepositoryImpl{Db: Db}
}

func (u UserRepositoryImpl) Save(user models.User) (uint, error) {
	result := u.Db.Create(&user)
	if err := result.Error; err != nil {
		return 0, errors.HandleDbError(err)
	}
	return user.ID, nil
}
