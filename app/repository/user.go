package repository

import (
	"futbook/app/dto"
	"futbook/app/model"
	"futbook/pkg/common"
	"futbook/platform/database"
)

type UserRepo struct {
	db *database.DB
}

func NewUserRepo(db *database.DB) UserRepository {
	return &UserRepo{db}
}

func (repo *UserRepo) GetByUsername(username string) (*model.User, error) {
	var user model.User
	result := repo.db.DB.First(&user, "phone = ?", username)
	return &user, result.Error
}

func (repo *UserRepo) Create(b *dto.CreateUser) error {
	hashed, _ := common.GeneratePasswordHash([]byte(b.Password))
	user := &model.User{
		FullName: b.FullName,
		Phone:    b.Username,
		Password: hashed,
		Role:     "user",
	}
	result := repo.db.DB.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
