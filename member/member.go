package member

import (
	"context"
	"strings"
)

type MemberStruct struct {
	Name     string `json:"name"`
	Position string `json:"position"`
	Email    string `json:"email"`
	Linkedin string `json:"linkedin"`
}

type MembersStruct struct {
	Members []MemberStruct
}

type ListResponse struct {
	Members []MemberStruct `json:"members"`
}

type SingleResponse struct {
	Member MemberStruct `json:"member"`
}

var members map[string]MemberStruct

func init() {
	// Call getMembers() once and cache result
	members = make(map[string]MemberStruct)
	for _, member := range getMembers() {
		members[strings.ToLower(member.Name)] = member
	}
}

// Returns all members sorted by name
// encore:api public path=/members method=GET
func Members(ctx context.Context) (*ListResponse, error) {
	msg := &ListResponse{Members: make([]MemberStruct, 0, len(members))}
	for _, member := range members {
		msg.Members = append(msg.Members, member)
	}

	// Sort members by name
	msg.Members = sortMembers(msg.Members)

	return msg, nil
}

// Returns a member
// encore:api public path=/member/:name method=GET
func Member(ctx context.Context, name string) (*SingleResponse, error) {

	// Lowercase name parameter
	name = strings.ToLower(name)

	// Lookup member directly in members map
	if member, ok := members[name]; ok {
		msg := &SingleResponse{Member: member}
		return msg, nil
	}

	// Return empty member if not found
	msg := &SingleResponse{Member: MemberStruct{}}
	return msg, nil
}
