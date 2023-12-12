package repository

import (
	"fiber-api-boilerplate/app/dto"
	"fiber-api-boilerplate/app/model"
)

type UserRepository interface {
	GetByUsername(username string) (*model.User, error)
	Create(b *dto.CreateUser) error
}

type RoleRepository interface {
}
