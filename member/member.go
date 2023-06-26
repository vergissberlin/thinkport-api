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
	msg := &Response{Members: "All Members "}
	return msg, nil
}
