package service

import (
	"context"
	"github.com/AntonioTrupac/hannaWebshop/model"
	service "github.com/AntonioTrupac/hannaWebshop/service/users"
	"github.com/AntonioTrupac/hannaWebshop/tools"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"gorm.io/gorm"
)

type AuthService interface {
	UserRegister(ctx context.Context, input *model.UserAuth) (interface{}, error)
	UserLogin(ctx context.Context, email string, password string) (interface{}, error)
	//UserRefresh(ctx context.Context, )
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

	token, err := tools.GenerateAccessToken(ctx, int(newUser.ID), input.Email)

	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"token": token,
		"email": input.Email,
	}, nil
}

// UserLogin implements AuthService
func (a *auth) UserLogin(ctx context.Context, email string, password string) (interface{}, error) {
	user, err := a.user.GetUserByEmail(ctx, email)
	if err != nil {
		// if user not found
		if err == gorm.ErrRecordNotFound {
			return nil, &gqlerror.Error{
				Message: "Email not found",
			}
		}

		return nil, err
	}

	if err := tools.ComparePassword(password, user.Password); err != nil {
		return nil, err
	}

	accessToken, err := tools.GenerateAccessToken(ctx, int(user.ID), user.Email)
	if err != nil {
		return nil, err
	}

	refreshToken, err := tools.GenerateRefreshToken(ctx, int(user.ID), user.Email)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
		"email":        user.Email,
	}, nil
}
