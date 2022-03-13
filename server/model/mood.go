package model

import "gorm.io/gorm"

type Mood struct {
	gorm.Model
	Name   string `gorm:"size:255;not null"`
	Notes  string `gorm:"size:255;not null"`
	UserId int
}
