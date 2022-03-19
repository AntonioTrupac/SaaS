package model

import "gorm.io/gorm"

type Mood struct {
	gorm.Model
	Notes      string
	UserId     int
	MoodTypeId int
	MoodType   MoodType
}
