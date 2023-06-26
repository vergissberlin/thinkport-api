package member

import (
	"context"
)

type Response struct {
	Members string `json:"members"`
}

// Returns all members
// encore:api public path=/members method=GET
func World(ctx context.Context, name string) (*Response, error) {
	msg := &Response{Members: "All Members "}
	return msg, nil
}
