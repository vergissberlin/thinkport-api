package member

import (
	"context"
)

type MemberStruct struct {
	Name     string
	Position string
	Email    string
	Linkedin string
}

type MembersStruct struct {
	Members []MemberStruct
}

type ListResponse struct {
	Members struct {
		Members []*MemberStruct `json:"members"`
	} `json:"members"`
}

type SingleResponse struct {
	Member MemberStruct `json:"MemberStruct"`
}

var members []MemberStruct

func init() {
	// Call getMembers() once and cache result
	members = getMembers()
}

// Returns all members
// encore:api public path=/members method=GET
func Members(ctx context.Context) (*ListResponse, error) {
	// Add a member manually to the slice
	members = append(members, MemberStruct{
		Name:     "John Doe",
		Position: "Software Engineer",
		Email:    "dirk@dirk.de",
		Linkedin: "https://www.linkedin.com/in/dirk/",
	})

	var memberPointers []*MemberStruct
	for i := range members {
		memberPointers = append(memberPointers, &members[i])
	}

	msg := &ListResponse{Members: struct {
		Members []*MemberStruct `json:"members"`
	}{Members: memberPointers}}

	return msg, nil
}

// Returns a member
// encore:api public path=/member/:name method=GET
func Member(ctx context.Context, name string) (*SingleResponse, error) {
	// Use map to lookup member by name
	memberMap := make(map[string]MemberStruct)
	for _, member := range members {
		memberMap[member.Name] = member
	}

	// Return member from map
	if member, ok := memberMap[name]; ok {
		msg := &SingleResponse{Member: member}

		return msg, nil
	}

	// Return empty member if not found
	msg := &SingleResponse{Member: MemberStruct{}}
	return msg, nil
}
