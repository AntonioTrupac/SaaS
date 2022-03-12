package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/AntonioTrupac/hannaWebshop/graph/generated"
)

func (r *mutationResolver) CreateMoods(ctx context.Context, input generated.MoodsInput) (*generated.Moods, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetMoods(ctx context.Context) ([]*generated.Moods, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetMoodByID(ctx context.Context, id int) (*generated.Moods, error) {
	panic(fmt.Errorf("not implemented"))
}
