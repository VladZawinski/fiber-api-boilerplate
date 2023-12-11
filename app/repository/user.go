package repository

import "fiber-api-boilerplate/platform/database"

type UserRepo struct {
	db *database.DB
}

func NewUserRepo(db *database.DB) UserRepository {
	return UserRepo{db}
}
