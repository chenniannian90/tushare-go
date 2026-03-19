package schema

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// APISchema represents the complete Tushare API schema
type APISchema struct {
	Version     string      `yaml:"version"`
	Description string      `yaml:"description"`
	Categories  []Category  `yaml:"categories"`
}

// Category represents a main API category (e.g., Stock, ETF, Index)
type Category struct {
	ID           string         `yaml:"id"`
	Name         string         `yaml:"name"`
	Subcategories []SubCategory  `yaml:"subcategories"`
}

// SubCategory represents a subcategory within a category
type SubCategory struct {
	ID   string `yaml:"id"`
	Name string `yaml:"name"`
	APIs []API  `yaml:"apis"`
}

// API represents a single API endpoint
type API struct {
	DocID   string `yaml:"doc_id"`
	Name    string `yaml:"name"`
	APIName string `yaml:"api_name"`
	URL     string `yaml:"url"`
}

// LoadSchema loads the API schema from YAML file
func LoadSchema() (*APISchema, error) {
	// Get the current directory
	dir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("failed to get working directory: %w", err)
	}

	// Try to find the schema file
	schemaPath := filepath.Join(dir, "schema", "api_schema.yaml")

	// If not found, try relative to the project root
	if _, err := os.Stat(schemaPath); os.IsNotExist(err) {
		schemaPath = filepath.Join(dir, "..", "schema", "api_schema.yaml")
	}

	data, err := os.ReadFile(schemaPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read schema file: %w", err)
	}

	var schema APISchema
	if err := yaml.Unmarshal(data, &schema); err != nil {
		return nil, fmt.Errorf("failed to parse schema YAML: %w", err)
	}

	return &schema, nil
}

// GetAPIByName finds an API by its api_name across all categories
func (s *APISchema) GetAPIByName(apiName string) (*API, error) {
	for _, cat := range s.Categories {
		for _, sub := range cat.Subcategories {
			for _, api := range sub.APIs {
				if api.APIName == apiName {
					return &api, nil
				}
			}
		}
	}
	return nil, fmt.Errorf("API not found: %s", apiName)
}

// GetAPIsByCategory returns all APIs in a specific category
func (s *APISchema) GetAPIsByCategory(categoryID string) ([]API, error) {
	for _, cat := range s.Categories {
		if cat.ID == categoryID {
			var apis []API
			for _, sub := range cat.Subcategories {
				apis = append(apis, sub.APIs...)
			}
			return apis, nil
		}
	}
	return nil, fmt.Errorf("category not found: %s", categoryID)
}

// GetAPIsBySubcategory returns all APIs in a specific subcategory
func (s *APISchema) GetAPIsBySubcategory(categoryID, subcategoryID string) ([]API, error) {
	for _, cat := range s.Categories {
		if cat.ID == categoryID {
			for _, sub := range cat.Subcategories {
				if sub.ID == subcategoryID {
					return sub.APIs, nil
				}
			}
			return nil, fmt.Errorf("subcategory not found: %s", subcategoryID)
		}
	}
	return nil, fmt.Errorf("category not found: %s", categoryID)
}

// TotalAPIs returns the total number of APIs in the schema
func (s *APISchema) TotalAPIs() int {
	total := 0
	for _, cat := range s.Categories {
		for _, sub := range cat.Subcategories {
			total += len(sub.APIs)
		}
	}
	return total
}

// ListCategories returns all category IDs and names
func (s *APISchema) ListCategories() []string {
	var categories []string
	for _, cat := range s.Categories {
		categories = append(categories, fmt.Sprintf("%s (%s)", cat.Name, cat.ID))
	}
	return categories
}

// ListSubcategories returns all subcategory IDs and names for a category
func (s *APISchema) ListSubcategories(categoryID string) ([]string, error) {
	for _, cat := range s.Categories {
		if cat.ID == categoryID {
			var subcategories []string
			for _, sub := range cat.Subcategories {
				subcategories = append(subcategories, fmt.Sprintf("%s (%s)", sub.Name, sub.ID))
			}
			return subcategories, nil
		}
	}
	return nil, fmt.Errorf("category not found: %s", categoryID)
}
