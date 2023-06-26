package member

import (
	"context"
	"testing"
)

// Overload init() to use test data
func init() {
	members = []MemberStruct{
		{Name: "John"},
		{Name: "Jane"},
	}
}

func TestMembers(t *testing.T) {
	expected := &ListResponse{
		Members: []MemberStruct{
			{Name: "John"},
			{Name: "Jane"},
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

func TestMember(t *testing.T) {
	expected := &SingleResponse{
		Member: MemberStruct{Name: "jane"},
	}

	result, err := Member(context.Background(), "JANE")

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if result.Member.Name != expected.Member.Name {
		t.Errorf("Expected member name %s, got %s", expected.Member.Name, result.Member.Name)
	}
}

func TestMemberNotFound(t *testing.T) {
	_, err := Member(context.Background(), "notfound")

	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}
