package about

// Override init() to use test data
func init() {
	// Create test data of 3 locations
	locations = make(map[string]LocationStruct)
	locations["a"] = LocationStruct{Name: "A"}
	locations["b"] = LocationStruct{Name: "B"}
}
