package types

import (
	"context"

	"github.com/AntonioTrupac/hannaWebshop/graph/generated"
	"github.com/AntonioTrupac/hannaWebshop/model"
	"github.com/AntonioTrupac/hannaWebshop/util"
)

//Create User mutation

//func ModelUser(ctx context.Context, in generated.UserInput) *model.User {
//	util.IsEmailValid(in.Email)
//
//	return &model.User{
//		Email:     in.Email,
//		FirstName: in.FirstName,
//		LastName:  in.LastName,
//		Age:       uint(in.Age),
//		Phone:     in.Phone,
//		Address:   mapAddressesToUser(in.Address),
//	}
//}

func mapAddressesToUser(addressInput []*generated.AddressInput) []*model.Address {
	var addresses []*model.Address

	for _, input := range addressInput {
		addresses = append(addresses, &model.Address{
			AddressLine: input.AddressLine,
			City:        input.City,
			Country:     input.Country,
			PostalCode:  input.PostalCode,
		})
	}

	return addresses
}

//func UserPayload(user *model.User) *generated.User {
//	return &generated.User{
//		ID:        int(user.ID),
//		Email:     user.Email,
//		FirstName: user.FirstName,
//		LastName:  user.LastName,
//		Age:       int(user.Age),
//		Phone:     user.Phone,
//		Address:   mapAddressPayloadToUser(user.Address),
//	}
//}

func mapAddressPayloadToUser(addressInput []*model.Address) []*generated.Address {
	var addresses []*generated.Address

	for _, addressInput := range addressInput {
		addresses = append(addresses, &generated.Address{
			PostalCode:  addressInput.PostalCode,
			Country:     addressInput.Country,
			City:        addressInput.City,
			AddressLine: addressInput.AddressLine,
		})
	}

	return addresses
}

// Users query
//func Users(users []*model.User) []*generated.User {
//	var out []*generated.User
//
//	for _, u := range users {
//		out = append(out, &generated.User{
//			FirstName: u.FirstName,
//			LastName:  u.LastName,
//			Age:       int(u.Age),
//			ID:        int(u.ID),
//			Email:     u.Email,
//			Phone:     u.Phone,
//			Moods:     MoodsPayload(u.Mood),
//		})
//	}
//
//	return out
//}

// ModelUserAuth Register mutation type
func ModelUserAuth(ctx context.Context, in generated.NewAuthUser) *model.UserAuth {
	util.IsEmailValid(in.Email)

	return &model.UserAuth{
		Email:     in.Email,
		FirstName: in.FirstName,
		LastName:  in.LastName,
		Password:  in.Password,
		Age:       uint(in.Age),
		Phone:     in.Phone,
		Address:   mapAddressesToUser(in.Address),
	}
}
