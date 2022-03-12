package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/AntonioTrupac/hannaWebshop/graph/generated"
)

func (r *mutationResolver) CreateHabits(ctx context.Context, input generated.HabitsInput) (*generated.Habits, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetHabits(ctx context.Context) ([]*generated.Habits, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetHabitsByID(ctx context.Context, id int) (*generated.Habits, error) {
	panic(fmt.Errorf("not implemented"))
}
