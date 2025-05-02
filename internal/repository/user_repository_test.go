package repository_test

import (
	"testing"

	"wallet-app/internal/models"
	"wallet-app/internal/utils"

	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Save(t *testing.T) {
	db, _, _, userRepo := utils.SetupTestDB()
	user := models.User{Name: "Test User"}
	userId, err := userRepo.Save(user)
	assert.NoError(t, err)
	assert.NotZero(t, userId)
	var savedUser models.User
	err = db.First(&savedUser, userId).Error
	assert.NoError(t, err)
	assert.Equal(t, user.Name, savedUser.Name)
}
