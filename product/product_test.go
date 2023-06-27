package product

import (
	"context"
	"testing"
)

// Override init() to use test data
func init() {
	// Create test data of 3 trainings
	trainings = make(map[string]TrainingStruct)
	trainings["a"] = TrainingStruct{Name: "A"}
	trainings["b"] = TrainingStruct{Name: "B"}
	trainings["c"] = TrainingStruct{Name: "C"}
}

func TestTrainings(t *testing.T) {

	expected := []TrainingStruct{
		{Name: "A"},
		{Name: "B"},
		{Name: "C"},
	}

	result, _ := Trainings(context.Background())

	if len(result.Trainings) != len(expected) {
		t.Errorf("Expected %d trainings, got %d", len(expected), len(result.Trainings))
	}

	for i := range result.Trainings {
		if result.Trainings[i].Name != expected[i].Name {
			t.Errorf("Expected training %d to be %s, got %s", i, expected[i].Name, result.Trainings[i].Name)
		}
	}
}
