package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/AntonioTrupac/hannaWebshop/graph/generated"
	"github.com/AntonioTrupac/hannaWebshop/graph/types"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input generated.UserInput) (*generated.User, error) {
	user := types.ModelUser(ctx, input)

	err := r.users.CreateAUser(user)

	if err != nil {
		return nil, err
	}

	return types.UserPayload(user), nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*generated.User, error) {
	users, err := r.users.GetUsers()

	if err != nil {
		return nil, err
	}

	return types.Users(users), nil
}