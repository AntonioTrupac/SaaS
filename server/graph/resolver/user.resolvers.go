package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/AntonioTrupac/hannaWebshop/graph/generated"
	"github.com/AntonioTrupac/hannaWebshop/graph/types"
)

func (r *authOpsResolver) Login(ctx context.Context, obj *generated.AuthOps, email string, password string) (interface{}, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *authOpsResolver) Register(ctx context.Context, obj *generated.AuthOps, input generated.NewAuthUser) (interface{}, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateUser(ctx context.Context, input generated.UserInput) (*generated.User, error) {
	user := types.ModelUser(ctx, input)

	err := r.users.CreateAUser(user)

	if err != nil {
		return nil, err
	}

	return types.UserPayload(user), nil
}

func (r *mutationResolver) Auth(ctx context.Context) (*generated.AuthOps, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Users(ctx context.Context) ([]*generated.User, error) {
	users, err := r.users.GetUsers()

	if err != nil {
		return nil, err
	}

	return types.Users(users), nil
}

func (r *queryResolver) GetUser(ctx context.Context, id int) (interface{}, error) {
	panic(fmt.Errorf("not implemented"))
}

// AuthOps returns generated.AuthOpsResolver implementation.
func (r *Resolver) AuthOps() generated.AuthOpsResolver { return &authOpsResolver{r} }

type authOpsResolver struct{ *Resolver }
