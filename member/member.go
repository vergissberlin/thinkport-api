package member

import (
	"context"
	"strings"
)

type MemberStruct struct {
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Position string `json:"position"`
	Details  string `json:"details"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
	Linkedin string `json:"linkedin"`
}

type MembersStruct struct {
	Members []MemberStruct
}

type MembersCountStruct struct {
	Count     int `json:"count"`
	Engineers int `json:"engineers"`
}

type ListResponse struct {
	Members []MemberStruct `json:"members"`
}

var members map[string]MemberStruct

func init() {
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

// Returns the count of members
// encore:api public path=/members/count method=GET
func MembersCount(ctx context.Context) (MembersCountStruct, error) {
	count := len(members)
	engineerCount := 0
	for _, member := range members {
		if strings.Contains(member.Position, "Engineer") {
			engineerCount += 1
		}
	}
	return MembersCountStruct{Count: count, Engineers: engineerCount}, nil
}

// Returns the count of engineers
// encore:api public path=/members/count/engineers method=GET
func EngineerCount(ctx context.Context) (int, error) {
	engineerCount := 0
	for _, member := range members {
		if strings.Contains(member.Position, "Engineer") {
			engineerCount += 1
		}
	}
	return engineerCount, nil
}

// Returns a MemberStruct
// encore:api public path=/member/:name method=GET
func Member(ctx context.Context, name string) (*MemberStruct, error) {
	// Find member by name
	member, ok := members[strings.ToLower(name)]
	if !ok {
		return nil, nil
	}

	return &member, nil
}

// Extract surname from email
// e.g. jdoe@ -> Doe
func extractSurnameFromEmail(email string) string {
	// Return empty string if email is empty
	if len(email) == 0 {
		return ""
	}
	// Extract string after  the first char before @
	surname := strings.Split(email, "@")[0]
	// Remove first char
	surname = surname[1:]

	// Capitalize first letter
	surname = strings.ToUpper(surname[0:1]) + surname[1:]

	return surname
}
