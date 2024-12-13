package product

import (
	"context"
	"testing"
)

func init() {
	trainings = map[string]TrainingStruct{
		"a": {Name: "A"},
		"b": {Name: "B"},
		"c": {Name: "C"},
	}
}

func TestTrainings(t *testing.T) {
	expected := []TrainingStruct{{Name: "A"}, {Name: "B"}, {Name: "C"}}
	result, _ := Trainings(context.Background())

	if len(result.Trainings) != len(expected) {
		t.Errorf("Expected %d trainings, got %d", len(expected), len(result.Trainings))
	}

	for i, training := range result.Trainings {
		if training.Name != expected[i].Name {
			t.Errorf("Expected training %d to be %s, got %s", i, expected[i].Name, training.Name)
		}
	}
}

func TestTrainingReturnsMatchingTrainings(t *testing.T) {
	expected := []TrainingStruct{{Name: "A"}}
	result, _ := Training(context.Background(), "a")

	if len(result.Trainings) != len(expected) {
		t.Errorf("Expected %d trainings, got %d", len(expected), len(result.Trainings))
	}

	for i, training := range result.Trainings {
		if training.Name != expected[i].Name {
			t.Errorf("Expected training %d to be %s, got %s", i, expected[i].Name, training.Name)
		}
	}
}

func TestTrainingReturnsEmptyWhenNoMatch(t *testing.T) {
	result, _ := Training(context.Background(), "z")

	if result != nil {
		t.Errorf("Expected nil, got %v", result)
	}
}

func TestTrainingHandlesCaseInsensitiveSearch(t *testing.T) {
	expected := []TrainingStruct{{Name: "A"}}
	result, _ := Training(context.Background(), "A")

	if len(result.Trainings) != len(expected) {
		t.Errorf("Expected %d trainings, got %d", len(expected), len(result.Trainings))
	}

	for i, training := range result.Trainings {
		if training.Name != expected[i].Name {
			t.Errorf("Expected training %d to be %s, got %s", i, expected[i].Name, training.Name)
		}
	}
}

func TestTrainingReturnsMultipleMatches(t *testing.T) {
	expected := []TrainingStruct{{Name: "A"}, {Name: "B"}}
	result, _ := Training(context.Background(), "a")

	if len(result.Trainings) != len(expected) {
		t.Errorf("Expected %d trainings, got %d", len(expected), len(result.Trainings))
	}

	for i, training := range result.Trainings {
		if training.Name != expected[i].Name {
			t.Errorf("Expected training %d to be %s, got %s", i, expected[i].Name, training.Name)
		}
	}
}
