package database_test

import (
	"encoding/json"
	"testing"

	"github.com/WindowsSov8forUs/sonolus-core-go/database"
)

func TestDatabasePostItemOmitsOptionalFields(t *testing.T) {
	item := database.DatabasePostItem{
		Name:    "post",
		Version: database.DatabasePostItemVersion,
		Title: database.LocalizationText{
			"en": "Post",
		},
		Author: database.LocalizationText{
			"en": "Author",
		},
		Tags: []database.DatabaseTag{},
	}

	data, err := json.Marshal(item)
	if err != nil {
		t.Fatal(err)
	}

	if string(data) == "" {
		t.Fatal("expected non-empty JSON")
	}
	if contains := jsonContains(data, "description") || jsonContains(data, "thumbnail"); contains {
		t.Fatalf("optional fields were not omitted: %s", data)
	}
}

func jsonContains(data []byte, field string) bool {
	var fields map[string]any
	if err := json.Unmarshal(data, &fields); err != nil {
		return false
	}
	_, ok := fields[field]
	return ok
}
