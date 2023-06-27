package product

import (
	"context"
	"strings"
)

type TrainingStruct struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Link        string `json:"link"`
}

type TrainingsStruct struct {
	Trainings []TrainingStruct
}

type TrainingListResponse struct {
	Trainings []TrainingStruct `json:"trainings"`
}

type TrainingSingleResponse struct {
	Training TrainingStruct `json:"training"`
}

var trainings map[string]TrainingStruct

func init() {
	// Call getTrainings() once and cache result
	trainings = make(map[string]TrainingStruct)
	for _, training := range getTrainings() {
		trainings[strings.ToLower(training.Name)] = training
	}

}

// Returns all trainings sorted by name
// encore:api public path=/product/trainings method=GET
func Trainings(ctx context.Context) *TrainingListResponse {
	msg := &TrainingListResponse{Trainings: make([]TrainingStruct, 0, len(trainings))}
	for _, training := range trainings {
		msg.Trainings = append(msg.Trainings, training)
	}

	// Sort trainings by name
	msg.Trainings = sortTrainings(msg.Trainings)

	return msg
}

// Returns a training
// encore:api public path=/product/training/:name method=GET
func Training(ctx context.Context, name string) (*TrainingListResponse, error) {

	// Lowercase name parameter
	name = strings.ToLower(name)

	// Filter trainings by name that contains name parameter
	var filteredTrainings []TrainingStruct
	for _, training := range trainings {
		if strings.Contains(strings.ToLower(training.Name), name) {
			filteredTrainings = append(filteredTrainings, training)
		}
	}

	// Return nil if no trainings found
	if len(filteredTrainings) == 0 {
		return nil, nil
	}

	// Return all trainings that match name parameter
	return &TrainingListResponse{Trainings: filteredTrainings}, nil

}
