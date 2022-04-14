package model

import "gorm.io/gorm"

type UserAuth struct {
	gorm.Model
	FirstName string `gorm:"size:255;not null"`
	LastName  string `gorm:"size:255;not null"`
	Email     string `gorm:"size:255;not null"`
	Password  string `gorm:"size:255;not null"`
	Age       uint
	Phone     string `gorm:"size:50;not null"`
	Address   []*Address
}
