package utils

import (
	"encoding/json"
	"gorm.io/datatypes"
)

// InterfaceToDataTypeJson converts a map[string]interface{} to datatypes.JSON
func InterfaceToDataTypeJson(dimensions map[string]interface{}) datatypes.JSON {
	if dimensions == nil {
		return datatypes.JSON("null")
	}

	// Marshal the map to JSON bytes
	bytes, err := json.Marshal(dimensions)
	if err != nil {
		// Return empty JSON object if marshaling fails
		return datatypes.JSON("{}")
	}

	// Convert to datatypes.JSON
	return datatypes.JSON(bytes)
}
