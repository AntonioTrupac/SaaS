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
	return r.auth.UserLogin(ctx, email, password)
}

func (r *authOpsResolver) Register(ctx context.Context, obj *generated.AuthOps, input generated.NewAuthUser) (interface{}, error) {
	return r.auth.UserRegister(ctx, types.ModelUserAuth(ctx, input))
}

func (r *mutationResolver) Auth(ctx context.Context) (*generated.AuthOps, error) {
	return &generated.AuthOps{}, nil
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input generated.RefreshTokenInput) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetUser(ctx context.Context, id int) (interface{}, error) {
	panic(fmt.Errorf("not implemented"))
}

// AuthOps returns generated.AuthOpsResolver implementation.
func (r *Resolver) AuthOps() generated.AuthOpsResolver { return &authOpsResolver{r} }

type authOpsResolver struct{ *Resolver }
