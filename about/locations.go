package about

import (
	"context"
	"strings"
)

// LocationStruct is a struct containing information about a location
type LocationStruct struct {
	Name    string `json:"name"`
	Comment string `json:"comment"`
	Address string `json:"address"`
}

type LocationListResponse struct {
	Locations []LocationStruct `json:"locations"`
}

var locations map[string]LocationStruct

func init() {
	// Call getLocations() once and cache result
	locations = make(map[string]LocationStruct)
	for _, location := range getLocations() {
		locations[strings.ToLower(location.Name)] = location
	}
}

// Endpoint to get information about the locations of the company
// encore:api public path=/about/locations method=GET
func Locations(ctx context.Context) (*LocationListResponse, error) {
	msg := &LocationListResponse{Locations: make([]LocationStruct, 0, len(locations))}
	for _, location := range locations {
		msg.Locations = append(msg.Locations, location)
	}

	return msg, nil
}
