package resolver

import (
	authService "github.com/AntonioTrupac/hannaWebshop/service/auth"
	moodsService "github.com/AntonioTrupac/hannaWebshop/service/moods"
	productsService "github.com/AntonioTrupac/hannaWebshop/service/products"
	usersService "github.com/AntonioTrupac/hannaWebshop/service/users"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	users    usersService.UserService
	products productsService.ProductService
	moods    moodsService.MoodsService
	auth     authService.AuthService
}

func NewResolver(users usersService.UserService, products productsService.ProductService, moods moodsService.MoodsService, auth authService.AuthService) *Resolver {
	return &Resolver{
		users:    users,
		products: products,
		moods:    moods,
		auth:     auth,
	}
}
