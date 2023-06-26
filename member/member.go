package member

import (
	"context"
)

type MemberStruct struct {
	Name     string `json:"name"`
	Position string `json:"position"`
	Email    string `json:"email"`
	Linkedin string `json:"linkedin"`
}

type MembersStruct struct {
	Members []MemberStruct `json:"members"`
}

type ListResponse struct {
	Members []MemberStruct `json:"members"`
}

type SingleResponse struct {
	Member MemberStruct `json:"member"`
}

var members []MemberStruct

func init() {
	// Call getMembers() once and cache result
	members = getMembers()
}

// Returns all members
// encore:api public path=/members method=GET
func Members(ctx context.Context) (*ListResponse, error) {

	var memberPointers []*MemberStruct
	for i := range members {
		memberPointers = append(memberPointers, &members[i])
	}

	msg := &ListResponse{Members: members}
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
