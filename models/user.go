package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Password string
	Email    string
	IsAdmin  bool
	Avatar   string
}

func (u *User) TableName() string {
	return "t_user"
}
