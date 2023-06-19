package hello

import (
	"context"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Response struct {
	Message string `json:"message"`
}

// World returns a greeting.
// encore:api public path=/hello/:name method=GET
func World(ctx context.Context, name string) (*Response, error) {
	msg := &Response{Message: "Hello " + name}
	return msg, nil
}

// Mars returns a greeting. It is a public API.
// encore:api public path=/mars/:name method=GET
func Mars(ctx context.Context, name string) (*Response, error) {
	msg := &Response{Message: "Hello from Mars " + name + "!"}
	return msg, nil
}

// Jupiter returns a greeting.
// encore:api public path=/jupiter/:name method=GET
func Jupiter(ctx context.Context, name string) (*Response, error) {
	// First letter of the name is capitalized
	name = cases.Title(language.English).String(name)
	msg := &Response{Message: "Hello from Jupiter " + name + "!!!"}
	return msg, nil
}
