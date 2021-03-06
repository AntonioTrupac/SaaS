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
			ID:         int(values.ID),
			Notes:      values.Notes,
			UserID:     values.UserId,
			MoodTypeID: values.MoodTypeId,
		})
	}

	return moods
}

// GetMoodById query

func MoodByIdPayload(moodPayload *model.Mood) *generated.Moods {
	return &generated.Moods{
		ID:         int(moodPayload.ID),
		Notes:      moodPayload.Notes,
		UserID:     moodPayload.UserId,
		MoodTypeID: moodPayload.MoodTypeId,
	}
}

// Create Mood mutation

func MapFromGeneratedInput(input generated.MoodsInput) *model.Mood {
	return &model.Mood{
		Notes: input.Notes,
	}
}

func GetPayloadFromDB(moodPayload *model.Mood) *generated.Moods {
	return &generated.Moods{
		ID:         int(moodPayload.ID),
		Notes:      moodPayload.Notes,
		UserID:     moodPayload.UserId,
		MoodTypeID: moodPayload.MoodTypeId,
	}
}
