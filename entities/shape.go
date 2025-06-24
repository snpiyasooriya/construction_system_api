package entities

type Shape struct {
	ID         uint
	Name       string
	Path       string
	Dimensions map[string]interface{} // Using a map to hold JSON-like data
	Labels     map[string]interface{} // Using a map to hold JSON-like data
}
