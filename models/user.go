package models

import (
	"database/sql/driver"
)

type BitBool bool

func (b BitBool) Value() (driver.Value, error) {
	if b {
		return byte(1), nil
	}
	return byte(0), nil
}

func (b *BitBool) Scan(v interface{}) error {
	*b = v.([]byte)[0] != 0
	return nil
}

type User struct {
	Model
	Username string  `json:"username"`
	Password string  `json:"password"`
	Email    string  `json:"eamil"`
	IsAdmin  BitBool `json:"is_admin"`
	Avatar   string  `json:"avatar"`
}

func (u *User) TableName() string {
	return "t_user"
}
