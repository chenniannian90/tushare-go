package schema_test

import (
	"fmt"
	"testing"

	"github.com/chenniannian90/tushare-go/schema"
)

func TestLoadSchema(t *testing.T) {
	s, err := schema.LoadSchema()
	if err != nil {
		t.Fatalf("Failed to load schema: %v", err)
	}

	// Test basic properties
	if s.Version == "" {
		t.Error("Version is empty")
	}

	if s.Description == "" {
		t.Error("Description is empty")
	}

	// Test categories
	if len(s.Categories) == 0 {
		t.Error("No categories found")
	}

	fmt.Printf("✅ Schema loaded successfully\n")
	fmt.Printf("Version: %s\n", s.Version)
	fmt.Printf("Description: %s\n", s.Description)
	fmt.Printf("Total APIs: %d\n", s.TotalAPIs())
}

func TestListCategories(t *testing.T) {
	s, err := schema.LoadSchema()
	if err != nil {
		t.Fatalf("Failed to load schema: %v", err)
	}

	categories := s.ListCategories()
	if len(categories) == 0 {
		t.Error("No categories returned")
	}

	fmt.Printf("\n📋 Categories (%d):\n", len(categories))
	for _, cat := range categories {
		fmt.Printf("  - %s\n", cat)
	}
}

func TestGetAPIByName(t *testing.T) {
	s, err := schema.LoadSchema()
	if err != nil {
		t.Fatalf("Failed to load schema: %v", err)
	}

	tests := []string{
		"stock_basic",
		"daily",
		"income",
		"trade_cal",
		"top10holders",
	}

	fmt.Printf("\n🔍 Looking up APIs:\n")
	for _, apiName := range tests {
		api, err := s.GetAPIByName(apiName)
		if err != nil {
			t.Errorf("Failed to find API %s: %v", apiName, err)
			continue
		}
		fmt.Printf("  ✅ %s: %s\n", api.APIName, api.Name)
	}
}

func TestGetAPIsByCategory(t *testing.T) {
	s, err := schema.LoadSchema()
	if err != nil {
		t.Fatalf("Failed to load schema: %v", err)
	}

	categories := []string{"stock", "etf", "index", "fund"}

	fmt.Printf("\n📊 APIs by Category:\n")
	for _, catID := range categories {
		apis, err := s.GetAPIsByCategory(catID)
		if err != nil {
			t.Errorf("Failed to get APIs for category %s: %v", catID, err)
			continue
		}
		fmt.Printf("  %s: %d APIs\n", catID, len(apis))
	}
}

func TestGetAPIsBySubcategory(t *testing.T) {
	s, err := schema.LoadSchema()
	if err != nil {
		t.Fatalf("Failed to load schema: %v", err)
	}

	// Test stock.basic subcategory
	apis, err := s.GetAPIsBySubcategory("stock", "basic")
	if err != nil {
		t.Errorf("Failed to get APIs for stock.basic: %v", err)
		return
	}

	fmt.Printf("\n📂 Stock - Basic APIs (%d):\n", len(apis))
	for _, api := range apis {
		fmt.Printf("  - %s: %s\n", api.APIName, api.Name)
	}
}

func ExampleLoadSchema() {
	s, err := schema.LoadSchema()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Tushare API Schema v%s\n", s.Version)
	fmt.Printf("Total APIs: %d\n", s.TotalAPIs())

	// List first 5 categories
	categories := s.ListCategories()
	for i, cat := range categories {
		if i >= 5 {
			break
		}
		fmt.Println(cat)
	}
}
