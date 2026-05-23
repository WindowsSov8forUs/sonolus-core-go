package core_test

import (
	"encoding/json"
	"errors"
	"testing"

	sonocore "github.com/WindowsSov8forUs/sonolus-core-go/core"
	sonerrors "github.com/WindowsSov8forUs/sonolus-core-go/errors"
)

func TestAutoExitUnmarshalJSON(t *testing.T) {
	var value sonocore.AutoExit
	if err := json.Unmarshal([]byte(`"fullCombo"`), &value); err != nil {
		t.Fatal(err)
	}
	if value != sonocore.AutoExitFullCombo {
		t.Fatalf("AutoExit = %q, want %q", value, sonocore.AutoExitFullCombo)
	}
}

func TestAutoExitUnmarshalJSONInvalidValue(t *testing.T) {
	var value sonocore.AutoExit
	err := json.Unmarshal([]byte(`"never"`), &value)
	if err == nil {
		t.Fatal("expected error")
	}

	var invalid sonerrors.InvalidEnumValueError
	if !errors.As(err, &invalid) {
		t.Fatalf("error = %T, want InvalidEnumValueError", err)
	}
	if invalid.TypeName != "AutoExit" || invalid.Value != "never" {
		t.Fatalf("invalid enum error = %#v", invalid)
	}
}

func TestItemTypeUnmarshalJSON(t *testing.T) {
	var value sonocore.ItemType
	if err := json.Unmarshal([]byte(`"level"`), &value); err != nil {
		t.Fatal(err)
	}
	if value != sonocore.ItemTypeLevel {
		t.Fatalf("ItemType = %q, want %q", value, sonocore.ItemTypeLevel)
	}
}
