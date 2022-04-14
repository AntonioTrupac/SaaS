package service

import (
	"context"

	"github.com/AntonioTrupac/hannaWebshop/model"
	service "github.com/AntonioTrupac/hannaWebshop/service/users"
	"github.com/AntonioTrupac/hannaWebshop/tools"
	"gorm.io/gorm"
)

type AuthService interface {
	UserRegister(ctx context.Context, input *model.UserAuth) (interface{}, error)
}

type auth struct {
	DB   *gorm.DB
	user service.UserService
}

func NewAuth(db *gorm.DB, user service.UserService) AuthService {
	return &auth{
		DB:   db,
		user: user,
	}
}

// UserRegister implements AuthService
func (a *auth) UserRegister(ctx context.Context, input *model.UserAuth) (interface{}, error) {
	//check if email exists and get the user if it does not
	_, err := a.user.GetUserByEmail(ctx, input.Password)

	if err == nil {
		if err != gorm.ErrRecordNotFound {
			return nil, err
		}
	}

	newUser, err := a.user.UserCreate(ctx, input)
	if err != nil {
		return nil, err
	}

	token, err := tools.GenerateJwt(ctx, int(newUser.ID), input.Email)

	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"token": token,
		"email": input.Email,
	}, nil
}
