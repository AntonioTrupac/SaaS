package service

import (
	"github.com/AntonioTrupac/hannaWebshop/model"
	"gorm.io/gorm"
)

type MoodsService interface {
	GetAllMoods() ([]*model.Mood, error)
	GetMoodByIdService(id int) (*model.Mood, error)
	CreateAMood(input *model.Mood) error
	CreateMoodTypes(input *model.MoodType) error
	GetAllMoodTypes() ([]*model.MoodType, error)
}

type Moods struct {
	DB *gorm.DB
}

func (m Moods) GetAllMoods() ([]*model.Mood, error) {
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

func (m Moods) CreateMoodTypes(input *model.MoodType) error {
	if err := m.DB.Create(input).Error; err != nil {
		return err
	}

	return nil
}

func (m Moods) GetAllMoodTypes() ([]*model.MoodType, error) {
	var moodTypes []*model.MoodType

	if err := m.DB.Find(&moodTypes).Error; err != nil {
		return nil, err
	}

	return moodTypes, nil
}

func NewMoods(db *gorm.DB) MoodsService {
	return &Moods{
		DB: db,
	}
}
