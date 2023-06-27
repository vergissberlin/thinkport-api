package product

import "testing"

func TestSortTrainings(t *testing.T) {
	trainings := []TrainingStruct{
		{Name: "C"},
		{Name: "B"},
		{Name: "A"},
	}

	expected := []TrainingStruct{
		{Name: "A"},
		{Name: "B"},
		{Name: "C"},
	}

	result := sortTrainings(trainings)

	if len(result) != len(expected) {
		t.Errorf("Expected %d trainings, got %d", len(expected), len(result))
	}

	for i := range result {
		if result[i].Name != expected[i].Name {
			t.Errorf("Expected training %d to be %s, got %s", i, expected[i].Name, result[i].Name)
		}
	}
}
