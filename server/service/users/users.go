package service

import (
	"context"
	"fmt"
	"github.com/AntonioTrupac/hannaWebshop/model"
	"github.com/AntonioTrupac/hannaWebshop/tools"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserService interface {
	GetUsers() ([]*model.User, error)
	CreateAUser(input *model.User) error
	UserCreate(ctx context.Context, input *model.UserAuth) (*model.UserAuth, error)
	GetUserByEmail(ctx context.Context, email string) (*model.UserAuth, error)
}

type users struct {
	DB *gorm.DB
}

func NewUsers(db *gorm.DB) UserService {
	return &users{
		DB: db,
	}
}

func (u *users) GetUsers() ([]*model.User, error) {
	var users []*model.User

	if err := u.DB.Find(&users).Error; err != nil {
		return nil, fmt.Errorf("/getUsers: Error getting users from the database! %v", err)
	}

	return users, nil
}

func (u *users) CreateAUser(input *model.User) error {
	return u.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Omit(clause.Associations).Create(input).Error; err != nil {
			return err
		}

		for _, value := range input.Address {
			value.UserId = int(input.ID)
		}

		if err := tx.CreateInBatches(input.Address, 100).Error; err != nil {
			return err
		}

		return nil
	})
}

func (u *users) UserCreate(ctx context.Context, input *model.UserAuth) (*model.UserAuth, error) {
	err := u.DB.Transaction(func(tx *gorm.DB) error {
		input.Password = tools.HashPassword(input.Password)
		if err := tx.Omit(clause.Associations).Create(input).Error; err != nil {
			return err
		}

		for _, value := range input.Address {
			value.UserAuthId = int(input.ID)
		}

		if err := tx.CreateInBatches(input.Address, 100).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}
	return input, nil
}

func (u *users) GetUserByEmail(ctx context.Context, email string) (*model.UserAuth, error) {
	var userAuth model.UserAuth

	if err := u.DB.Model(userAuth).Where("email LIKE ?", email).Take(&userAuth).Error; err != nil {
		return nil, err
	}

	return &userAuth, nil
}
