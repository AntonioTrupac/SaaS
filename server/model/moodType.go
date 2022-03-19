package model

import "gorm.io/gorm"

type MoodType struct {
	gorm.Model
	Name string
}
