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

func TestMembersReturnsAllMembers(t *testing.T) {
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

func TestMembersCountReturnsCorrectCounts(t *testing.T) {
	expected := MembersCountStruct{
		Count:      2,
		Architects: 0,
		Developers: 0,
		Engineers:  0,
		Managers:   0,
	}

	result, err := MembersCount(context.Background())

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestMemberReturnsCorrectMember(t *testing.T) {
	expected := &MemberStruct{Name: "John"}
	result, err := Member(context.Background(), "John")

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if result == nil || result.Name != expected.Name {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestMemberReturnsNilForNonExistentMember(t *testing.T) {
	result, err := Member(context.Background(), "NonExistent")

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if result != nil {
		t.Errorf("Expected nil, got %v", result)
	}
}
