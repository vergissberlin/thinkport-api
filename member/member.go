package member

import (
	"context"
	"strings"
)

type MemberStruct struct {
	Hash     string `json:"hash"`
	Name     string `json:"name"`
	Position string `json:"position"`
	Avatar   string `json:"avatar"`
	Linkedin string `json:"linkedin"`
}

type MembersStruct struct {
	Members []MemberStruct `json:"members"`
}

type MembersCountStruct struct {
	Count      int `json:"count"`
	Architects int `json:"architects"`
	Developers int `json:"developers"`
	Engineers  int `json:"engineers"`
	Managers   int `json:"managers"`
}

type ListResponse struct {
	Members []MemberStruct `json:"members"`
}

var members map[string]MemberStruct

func init() {
	members = make(map[string]MemberStruct)
	for _, member := range getMembers() {
		// generate hash of name and position
		// to avoid duplicates
		hash := hashMember(member)
		members[hash] = member

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
	architectCount := 0
	developerCount := 0
	engineerCount := 0
	managerCount := 0
	for _, member := range members {
		if strings.Contains(member.Position, "Architect") {
			architectCount += 1
		}
		if strings.Contains(member.Position, "Developer") {
			developerCount += 1
		}
		if strings.Contains(member.Position, "Engineer") {
			engineerCount += 1
		}
		if strings.Contains(member.Position, "Manager") || strings.Contains(member.Position, "CEO") {
			managerCount += 1
		}
	}
	return MembersCountStruct{
		Count:      count,
		Architects: architectCount,
		Developers: developerCount,
		Engineers:  engineerCount,
		Managers:   managerCount,
	}, nil
}

// Returns a MemberStruct
// encore:api public path=/member/:hash method=GET
func Member(ctx context.Context, hash string) (*MemberStruct, error) {
	member, ok := members[hash]
	if !ok {
		return nil, nil
	}

	return &member, nil
}
