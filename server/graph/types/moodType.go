package types

import (
	"github.com/AntonioTrupac/hannaWebshop/graph/generated"
	"github.com/AntonioTrupac/hannaWebshop/model"
)

// Create types mutations

func MapFromGeneratedTypeInput(input *generated.MoodTypeInput) *model.MoodType {
	return &model.MoodType{
		Name: input.Name,
	}
}

func MoodTypePayload(payload *model.MoodType) *generated.MoodType {
	return &generated.MoodType{
		ID:   int(payload.ID),
		Name: payload.Name,
	}
}

// get mood types query

func GetMoodTypesPayload(payload []*model.MoodType) []*generated.MoodType {
	var moodPayload []*generated.MoodType

	for _, values := range payload {
		moodPayload = append(moodPayload, &generated.MoodType{
			ID:   int(values.ID),
			Name: values.Name,
		})
	}

	return moodPayload
}
