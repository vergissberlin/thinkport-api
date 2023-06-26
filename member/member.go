package member

import (
	"context"
)

type Response struct {
	Members string `json:"members"`
}

// Returns all members
// encore:api public path=/members method=GET
func Members(ctx context.Context) (*Response, error) {
	var members string = crawlMembers()
	msg := &Response{Members: members}
	return msg, nil
}
