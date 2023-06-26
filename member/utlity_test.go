package member

import "testing"

func TestSortMembers(t *testing.T) {
    members := []MemberStruct{
        {Name: "c"},
        {Name: "b"},
        {Name: "a"},
    }

    expected := []MemberStruct{
        {Name: "a"},
        {Name: "b"},
        {Name: "c"},
    }

    result := sortMembers(members)

    if len(result) != len(expected) {
        t.Errorf("Expected sorted slice of length %d, got %d", len(expected), len(result))
    }

    for i := range result {
        if result[i].Name != expected[i].Name {
            t.Errorf("Expected member at index %d to be %s, got %s", i, expected[i].Name, result[i].Name)
        }
    }
}
