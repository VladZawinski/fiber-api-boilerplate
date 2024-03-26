package repository

import (
	"futbook/app/dto"
	"futbook/app/model"
)

type UserRepository interface {
	GetByUsername(username string) (*model.User, error)
	Create(b *dto.CreateUser) error
}

type RoleRepository interface {
}
