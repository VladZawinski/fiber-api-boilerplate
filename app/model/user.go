package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FullName string `gorm:"column:full_name"`
	Password string `gorm:"column:password"`
	Phone    string `gorm:"column:phone"`
	Role     string `gorm:"column:role"`
}
