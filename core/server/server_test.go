package server_test

import (
	"testing"

	"github.com/WindowsSov8forUs/sonolus-core-go/core"
	"github.com/WindowsSov8forUs/sonolus-core-go/core/server"
)

func TestServerInfoItemButtonButtonType(t *testing.T) {
	button := server.ServerInfoItemButton{
		Type: core.ItemTypeLevel,
	}

	if got := button.ButtonType(); got != "level" {
		t.Fatalf("ButtonType() = %q, want %q", got, "level")
	}
}
