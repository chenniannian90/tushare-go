package gen

import (
	"encoding/json"
	"os"
)

// APISpec represents the API specification loaded from JSON
type APISpec struct {
	APIName        string        `json:"api_name"`
	Description    string        `json:"description"`
	RequestParams  []ParamField  `json:"request_params"`
	ResponseFields []ParamField  `json:"response_fields"`
}

// ParamField represents a parameter or field definition
type ParamField struct {
	Name        string       `json:"name"`
	Type        string       `json:"type"`
	Description string       `json:"description"`
	Properties  []ParamField `json:"properties,omitempty"`  // For nested object types
	Items       *ParamField  `json:"items,omitempty"`       // For array types
	Enum        []string     `json:"enum,omitempty"`        // For enum types
}

// LoadSpec loads an API specification from a JSON file
func LoadSpec(specFile string) (*APISpec, error) {
	data, err := os.ReadFile(specFile)
	if err != nil {
		return nil, err
	}

	var spec APISpec
	if err := json.Unmarshal(data, &spec); err != nil {
		return nil, err
	}

	return &spec, nil
}
