package types

import (
	"github.com/AntonioTrupac/hannaWebshop/graph/generated"
	"github.com/AntonioTrupac/hannaWebshop/model"
)

// GetMoods query

func MoodsPayload(moodPayload []*model.Mood) []*generated.Moods {
	var moods []*generated.Moods

	for _, values := range moodPayload {
		moods = append(moods, &generated.Moods{
			ID:     int(values.ID),
			Name:   values.Name,
			Notes:  values.Notes,
			UserID: values.UserId,
		})
	}

	return moods
}

// GetMoodById query

func MoodByIdPayload(moodPayload *model.Mood) *generated.Moods {
	return &generated.Moods{
		ID:     int(moodPayload.ID),
		Name:   moodPayload.Name,
		Notes:  moodPayload.Notes,
		UserID: moodPayload.UserId,
	}
}

// Create Mood mutation

func MapFromGeneratedInput(input generated.MoodsInput) *model.Mood {
	return &model.Mood{
		Name:  input.Name,
		Notes: input.Notes,
	}
}

func GetPayloadFromDB(moodPayload *model.Mood) *generated.Moods {
	return &generated.Moods{
		ID:     int(moodPayload.ID),
		Notes:  moodPayload.Notes,
		Name:   moodPayload.Name,
		UserID: moodPayload.UserId,
	}
}
