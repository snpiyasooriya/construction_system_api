package utils

import (
	"encoding/json"
	"gorm.io/datatypes"
)

func DataTypeJsonToInterface(data datatypes.JSON) map[string]interface{} {
	var dimensions map[string]interface{}
	// Use json.Unmarshal instead of s.Dimensions.Unmarshal
	if len(data) > 0 {
		if err := json.Unmarshal(data, &dimensions); err != nil {
			// If there's an error, initialize an empty map
			return make(map[string]interface{})
		}
	} else {
		return make(map[string]interface{})
	}
	return dimensions
}
