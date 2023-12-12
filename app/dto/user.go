package dto

import "fiber-api-boilerplate/app/model"

type User struct {
	ID        uint   `json:"id"`
	Phone     string `json:"phone"`
	FullName  string `json:"fullName"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	IsAdmin   bool   `json:"isAdmin"`
}

type Auth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
	FullName string `json:"fullName"`
}

func ToUser(u *model.User) *User {
	return &User{
		ID:        u.ID,
		CreatedAt: u.CreatedAt.String(),
		UpdatedAt: u.UpdatedAt.String(),
		IsAdmin:   u.Role == "admin",
		Phone:     u.Phone,
		FullName:  u.FullName,
	}
}
