package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/AntonioTrupac/hannaWebshop/graph/generated"
	"github.com/AntonioTrupac/hannaWebshop/graph/types"
)

func (r *mutationResolver) CreateMoods(ctx context.Context, input generated.MoodsInput) (*generated.Moods, error) {
	mood := types.MapFromGeneratedInput(input)

	err := r.moods.CreateAMood(mood)

	if err != nil {
		return nil, err
	}

	return types.GetPayloadFromDB(mood), nil
}

func (r *mutationResolver) CreateMoodTypes(ctx context.Context, input *generated.MoodTypeInput) (*generated.MoodType, error) {
	moodType := types.MapFromGeneratedTypeInput(input)

	err := r.moods.CreateMoodTypes(moodType)

	if err != nil {
		return nil, err
	}

	return types.MoodTypePayload(moodType), nil
}

func (r *queryResolver) GetMoods(ctx context.Context) ([]*generated.Moods, error) {
	moods, err := r.moods.GetAllMoods()

	if err != nil {
		return nil, err
	}

	return types.MoodsPayload(moods), nil
}

func (r *queryResolver) GetMoodByID(ctx context.Context, id int) (*generated.Moods, error) {
	mood, err := r.moods.GetMoodByIdService(id)

	if err != nil {
		return nil, err
	}

	return types.MoodByIdPayload(mood), nil
}

func (r *queryResolver) GetMoodTypes(ctx context.Context) ([]*generated.MoodType, error) {
	moods, err := r.moods.GetAllMoodTypes()

	if err != nil {
		return nil, err
	}

	return types.GetMoodTypesPayload(moods), nil
}
