package core

import (
	"encoding/json"
	"testing"

	"github.com/WindowsSov8forUs/sonolus-core-go/internal/nilable"
)

func TestSrlMarshalOmitMissingFields(t *testing.T) {
	data, err := json.Marshal(Srl{})
	if err != nil {
		t.Fatal(err)
	}

	if got, want := string(data), `{}`; got != want {
		t.Fatalf("MarshalJSON() = %s, want %s", got, want)
	}
}

func TestSrlMarshalExplicitNull(t *testing.T) {
	hash := nilable.Nil[string]()
	data, err := json.Marshal(Srl{Hash: &hash})
	if err != nil {
		t.Fatal(err)
	}

	if got, want := string(data), `{"hash":null}`; got != want {
		t.Fatalf("MarshalJSON() = %s, want %s", got, want)
	}
}

func TestSrlUnmarshalValue(t *testing.T) {
	var srl Srl
	if err := json.Unmarshal([]byte(`{"hash":"abc"}`), &srl); err != nil {
		t.Fatal(err)
	}

	if srl.Hash == nil {
		t.Fatal("hash pointer is nil")
	}
	value, ok := srl.Hash.Value()
	if !ok || value != "abc" {
		t.Fatalf("hash = %q, %v; want %q, true", value, ok, "abc")
	}
}

func TestSrlUnmarshalExplicitNull(t *testing.T) {
	var srl Srl
	if err := json.Unmarshal([]byte(`{"hash":null}`), &srl); err != nil {
		t.Fatal(err)
	}
	if srl.Hash == nil || !srl.Hash.IsNil() || srl.URL != nil {
		t.Fatalf("Srl = %#v, want explicit null hash and absent url", srl)
	}
	data, err := json.Marshal(srl)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := string(data), `{"hash":null}`; got != want {
		t.Fatalf("round trip = %s, want %s", got, want)
	}
}
