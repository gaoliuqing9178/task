package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Account  string `json:"account" gorm:"unique"`
	Password string `json:"password"`
}

func (u *User) TableName() string {
	return "user"
}
