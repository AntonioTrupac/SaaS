package resolver

import (
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
}

func NewResolver(users usersService.UserService, products productsService.ProductService, moods moodsService.MoodsService) *Resolver {
	return &Resolver{
		users:    users,
		products: products,
		moods:    moods,
	}
}
