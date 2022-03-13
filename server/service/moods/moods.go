package service

import (
	"github.com/AntonioTrupac/hannaWebshop/model"
	"gorm.io/gorm"
)

type MoodsService interface {
	GetMoodsService() ([]*model.Mood, error)
	GetMoodByIdService(id int) (*model.Mood, error)
	CreateAMood(input *model.Mood) error
}

type Moods struct {
	DB *gorm.DB
}

func (m Moods) GetMoodsService() ([]*model.Mood, error) {
	var moods []*model.Mood

	if err := m.DB.Find(&moods).Error; err != nil {
		return nil, err
	}

	return moods, nil
}

func (m Moods) GetMoodByIdService(id int) (*model.Mood, error) {
	var mood *model.Mood

	if err := m.DB.Where("id = ?", id).Find(&mood).Error; err != nil {
		return nil, err
	}

	return mood, nil
}

func (m Moods) CreateAMood(input *model.Mood) error {
	return m.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(input).Error; err != nil {
			return err
		}

		return nil
	})
}

func NewMoods(db *gorm.DB) MoodsService {
	return &Moods{
		DB: db,
	}
}
