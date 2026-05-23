package server_test

import (
	"encoding/json"
	stderrors "errors"
	"reflect"
	"testing"

	"github.com/WindowsSov8forUs/sonolus-core-go/core"
	"github.com/WindowsSov8forUs/sonolus-core-go/core/resource"
	"github.com/WindowsSov8forUs/sonolus-core-go/core/server"
	coreerrors "github.com/WindowsSov8forUs/sonolus-core-go/errors"
)

func TestDecodeServerOption(t *testing.T) {
	tests := []struct {
		name string
		data string
		want any
	}{
		{"text", `{"type":"text"}`, server.ServerTextOption{}},
		{"textArea", `{"type":"textArea"}`, server.ServerTextAreaOption{}},
		{"slider", `{"type":"slider"}`, server.ServerSliderOption{}},
		{"toggle", `{"type":"toggle"}`, server.ServerToggleOption{}},
		{"select", `{"type":"select"}`, server.ServerSelectOption{}},
		{"multi", `{"type":"multi"}`, server.ServerMultiOption{}},
		{"serverItem", `{"type":"serverItem","itemType":"level"}`, server.ServerServerItemOption{}},
		{"serverItems", `{"type":"serverItems","itemType":"level"}`, server.ServerServerItemsOption{}},
		{"collectionItem", `{"type":"collectionItem","itemType":"level"}`, server.ServerCollectionItemOption{}},
		{"file", `{"type":"file"}`, server.ServerFileOption{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := server.DecodeServerOption([]byte(tt.data))
			if err != nil {
				t.Fatal(err)
			}
			if reflect.TypeOf(got) != reflect.TypeOf(tt.want) {
				t.Fatalf("DecodeServerOption() = %T, want %T", got, tt.want)
			}
		})
	}
}

func TestServerFormUnmarshalJSON(t *testing.T) {
	var form server.ServerForm
	if err := json.Unmarshal([]byte(`{
		"type":"search",
		"title":"#SEARCH",
		"icon":"search",
		"requireConfirmation":false,
		"options":[
			{"type":"text","query":"q"},
			{"type":"toggle","query":"ranked"}
		]
	}`), &form); err != nil {
		t.Fatal(err)
	}

	if _, ok := form.Options[0].(server.ServerTextOption); !ok {
		t.Fatalf("Options[0] = %T, want ServerTextOption", form.Options[0])
	}
	if _, ok := form.Options[1].(server.ServerToggleOption); !ok {
		t.Fatalf("Options[1] = %T, want ServerToggleOption", form.Options[1])
	}
}

func TestServerInfoUnmarshalJSON(t *testing.T) {
	var info server.ServerInfo
	if err := json.Unmarshal([]byte(`{
		"title":"title",
		"buttons":[
			{"type":"authentication"},
			{"type":"configuration"},
			{"type":"level","title":"#LEVEL"}
		],
		"configuration":{"options":[{"type":"text"}]}
	}`), &info); err != nil {
		t.Fatal(err)
	}

	if _, ok := info.Buttons[0].(server.ServerInfoAuthenticationButton); !ok {
		t.Fatalf("Buttons[0] = %T, want ServerInfoAuthenticationButton", info.Buttons[0])
	}
	if _, ok := info.Buttons[1].(server.ServerInfoConfigurationButton); !ok {
		t.Fatalf("Buttons[1] = %T, want ServerInfoConfigurationButton", info.Buttons[1])
	}
	if button, ok := info.Buttons[2].(server.ServerInfoItemButton); !ok {
		t.Fatalf("Buttons[2] = %T, want ServerInfoItemButton", info.Buttons[2])
	} else if button.Type != core.ItemTypeLevel {
		t.Fatalf("Buttons[2].Type = %q, want %q", button.Type, core.ItemTypeLevel)
	}
	if _, ok := info.Configuration.Options[0].(server.ServerTextOption); !ok {
		t.Fatalf("Configuration.Options[0] = %T, want ServerTextOption", info.Configuration.Options[0])
	}
}

func TestDecodeServerOptionUnknownType(t *testing.T) {
	_, err := server.DecodeServerOption([]byte(`{"type":"unknown"}`))
	var target coreerrors.UnknownUnionTypeError
	if !stderrors.As(err, &target) {
		t.Fatalf("error = %T, want UnknownUnionTypeError", err)
	}
	if target.Union != "ServerOption" || target.Type != "unknown" {
		t.Fatalf("error = %#v", target)
	}
}

func TestDecodeServerOptionInvalidJSON(t *testing.T) {
	if _, err := server.DecodeServerOption([]byte(`{`)); err == nil {
		t.Fatal("DecodeServerOption() error = nil, want JSON error")
	}
}

func TestDecodeMultiplayerUnions(t *testing.T) {
	command, err := server.DecodeCommand([]byte(`{
		"type":"addChatMessage",
		"nonce":1,
		"message":{"type":"quick","userId":"u1","value":"hello"}
	}`))
	if err != nil {
		t.Fatal(err)
	}
	addCommand, ok := command.(server.AddChatMessageCommand)
	if !ok {
		t.Fatalf("DecodeCommand() = %T, want AddChatMessageCommand", command)
	}
	if _, ok := addCommand.Message.(server.QuickChatMessage); !ok {
		t.Fatalf("Message = %T, want QuickChatMessage", addCommand.Message)
	}

	event, err := server.DecodeEvent([]byte(`{
		"type":"updateOptions",
		"optionValues":"",
		"options":[{
			"type":"configuration",
			"title":"#CONFIGURATION",
			"icon":"settings",
			"requireConfirmation":false,
			"options":[{"type":"slider","min":0,"max":1,"step":1}]
		}]
	}`))
	if err != nil {
		t.Fatal(err)
	}
	updateOptions, ok := event.(server.UpdateOptionsEvent)
	if !ok {
		t.Fatalf("DecodeEvent() = %T, want UpdateOptionsEvent", event)
	}
	if _, ok := updateOptions.Options[0].Options[0].(server.ServerSliderOption); !ok {
		t.Fatalf("nested option = %T, want ServerSliderOption", updateOptions.Options[0].Options[0])
	}
}

func TestServerItemSectionUnmarshalJSON(t *testing.T) {
	var section server.ServerItemSection
	if err := json.Unmarshal([]byte(`{
		"title":"title",
		"itemType":"level",
		"items":[
			{"name":"level-1","version":1},
			{"name":"level-2","version":1}
		]
	}`), &section); err != nil {
		t.Fatal(err)
	}

	if _, ok := section.Items[0].(resource.LevelItem); !ok {
		t.Fatalf("Items[0] = %T, want LevelItem", section.Items[0])
	}
	if _, ok := section.Items[1].(resource.LevelItem); !ok {
		t.Fatalf("Items[1] = %T, want LevelItem", section.Items[1])
	}
}
