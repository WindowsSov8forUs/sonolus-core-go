package database_test

import (
	"testing"

	"github.com/WindowsSov8forUs/sonolus-core-go/core"
	"github.com/WindowsSov8forUs/sonolus-core-go/database"
)

func TestLocalizePreferredLocale(t *testing.T) {
	text := database.LocalizationText{
		"en": core.Text("Title"),
		"ja": core.Text("タイトル"),
	}

	if got := database.Localize(text, "ja", "en"); got != "タイトル" {
		t.Fatalf("Localize() = %q", got)
	}
}

func TestLocalizeFallbackLocale(t *testing.T) {
	text := database.LocalizationText{
		"en": core.Text("Title"),
	}

	if got := database.Localize(text, "ja", "en"); got != "Title" {
		t.Fatalf("Localize() = %q", got)
	}
}

func TestLocalizeEmptyText(t *testing.T) {
	if got := database.Localize(nil, "ja", "en"); got != "" {
		t.Fatalf("Localize() = %q, want empty string", got)
	}
}
