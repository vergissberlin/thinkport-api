package member

import (
	"context"
	"testing"
)

// Overload init() to use test data
func init() {
	members = map[string]MemberStruct{
		"John": {Name: "John"},
		"Jane": {Name: "Jane"},
	}
}

func TestMembers(t *testing.T) {
	expected := &ListResponse{
		Members: []MemberStruct{
			{Name: "Jane"},
			{Name: "John"},
		},
	}

	result, err := Members(context.Background())

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if len(result.Members) != len(expected.Members) {
		t.Errorf("Expected %d members, got %d", len(expected.Members), len(result.Members))
	}

	for i := range result.Members {
		if result.Members[i].Name != expected.Members[i].Name {
			t.Errorf("Expected member %d to be %s, got %s", i, expected.Members[i].Name, result.Members[i].Name)
		}
	}
}
