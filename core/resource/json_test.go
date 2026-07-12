package resource_test

import (
	"encoding/json"
	stderrors "errors"
	"reflect"
	"strings"
	"testing"

	"github.com/WindowsSov8forUs/sonolus-core-go/core"
	"github.com/WindowsSov8forUs/sonolus-core-go/core/resource"
	coreerrors "github.com/WindowsSov8forUs/sonolus-core-go/errors"
)

func TestDecodeItem(t *testing.T) {
	tests := []struct {
		itemType core.ItemType
		want     any
	}{
		{core.ItemTypePost, resource.PostItem{}},
		{core.ItemTypePlaylist, resource.PlaylistItem{}},
		{core.ItemTypeLevel, resource.LevelItem{}},
		{core.ItemTypeSkin, resource.SkinItem{}},
		{core.ItemTypeBackground, resource.BackgroundItem{}},
		{core.ItemTypeEffect, resource.EffectItem{}},
		{core.ItemTypeParticle, resource.ParticleItem{}},
		{core.ItemTypeEngine, resource.EngineItem{}},
		{core.ItemTypeReplay, resource.ReplayItem{}},
		{core.ItemTypeRoom, resource.RoomItem{}},
		{core.ItemTypeUser, resource.UserItem{}},
	}

	for _, tt := range tests {
		t.Run(string(tt.itemType), func(t *testing.T) {
			got, err := resource.DecodeItem(tt.itemType, []byte(`{"name":"name"}`))
			if err != nil {
				t.Fatal(err)
			}
			if reflect.TypeOf(got) != reflect.TypeOf(tt.want) {
				t.Fatalf("DecodeItem() = %T, want %T", got, tt.want)
			}
		})
	}
}

func TestDecodeItemUnknownType(t *testing.T) {
	_, err := resource.DecodeItem(core.ItemType("unknown"), []byte(`{}`))
	var target coreerrors.UnknownUnionTypeError
	if !stderrors.As(err, &target) {
		t.Fatalf("error = %T, want UnknownUnionTypeError", err)
	}
	if target.Union != "ResourceItem" || target.Type != "unknown" {
		t.Fatalf("error = %#v", target)
	}
}

func TestEngineConfigurationUnmarshalJSON(t *testing.T) {
	var configuration resource.EngineConfiguration
	if err := json.Unmarshal([]byte(`{
		"options":[
			{"type":"slider"},
			{"type":"toggle"},
			{"type":"select"}
		],
		"ui":{}
	}`), &configuration); err != nil {
		t.Fatal(err)
	}

	if _, ok := configuration.Options[0].(resource.EngineConfigurationSliderOption); !ok {
		t.Fatalf("Options[0] = %T, want EngineConfigurationSliderOption", configuration.Options[0])
	}
	if _, ok := configuration.Options[1].(resource.EngineConfigurationToggleOption); !ok {
		t.Fatalf("Options[1] = %T, want EngineConfigurationToggleOption", configuration.Options[1])
	}
	if _, ok := configuration.Options[2].(resource.EngineConfigurationSelectOption); !ok {
		t.Fatalf("Options[2] = %T, want EngineConfigurationSelectOption", configuration.Options[2])
	}
}

func TestEngineConfigurationOptionTitleRoundTrip(t *testing.T) {
	option := resource.EngineConfigurationSliderOption{
		EngineConfigurationOptionBase: resource.EngineConfigurationOptionBase{
			Name:  "speed",
			Title: "Speed",
		},
		Type: resource.EngineConfigurationOptionTypeSlider,
	}
	data, err := json.Marshal(option)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(data), `"title":"Speed"`) {
		t.Fatalf("JSON = %s, want title", data)
	}

	decoded, err := resource.DecodeEngineConfigurationOption(data)
	if err != nil {
		t.Fatal(err)
	}
	decodedOption := decoded.(resource.EngineConfigurationSliderOption)
	if decodedOption.Title != "Speed" {
		t.Fatalf("Title = %q, want Speed", decodedOption.Title)
	}

	option.Title = ""
	data, err = json.Marshal(option)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(data), `"title"`) {
		t.Fatalf("JSON = %s, want omitted title", data)
	}
}

func TestEngineDataNodesUnmarshalJSON(t *testing.T) {
	var data resource.EnginePlayData
	if err := json.Unmarshal([]byte(`{
		"skin":{"sprites":[]},
		"effect":{"clips":[]},
		"particle":{"effects":[]},
		"buckets":[],
		"archetypes":[],
		"nodes":[
			{"value":1},
			{"func":"Add","args":[0,1]}
		]
	}`), &data); err != nil {
		t.Fatal(err)
	}

	if _, ok := data.Nodes[0].(resource.EngineDataValueNode); !ok {
		t.Fatalf("Nodes[0] = %T, want EngineDataValueNode", data.Nodes[0])
	}
	if _, ok := data.Nodes[1].(resource.EngineDataFunctionNode); !ok {
		t.Fatalf("Nodes[1] = %T, want EngineDataFunctionNode", data.Nodes[1])
	}
}

func TestLevelDataEntityUnmarshalJSON(t *testing.T) {
	var entity resource.LevelDataEntity
	if err := json.Unmarshal([]byte(`{
		"archetype":"#BPM_CHANGE",
		"data":[
			{"name":"#BEAT","value":1},
			{"name":"#BPM","ref":"bpm"}
		]
	}`), &entity); err != nil {
		t.Fatal(err)
	}

	if _, ok := entity.Data[0].(resource.LevelDataEntityValueData); !ok {
		t.Fatalf("Data[0] = %T, want LevelDataEntityValueData", entity.Data[0])
	}
	if _, ok := entity.Data[1].(resource.LevelDataEntityRefData); !ok {
		t.Fatalf("Data[1] = %T, want LevelDataEntityRefData", entity.Data[1])
	}
}

func TestDecodeEngineDataNodeUnknownShape(t *testing.T) {
	_, err := resource.DecodeEngineDataNode([]byte(`{"args":[]}`))
	var target coreerrors.UnknownUnionShapeError
	if !stderrors.As(err, &target) {
		t.Fatalf("error = %T, want UnknownUnionShapeError", err)
	}
	if target.Union != "EngineDataNode" {
		t.Fatalf("error = %#v", target)
	}
}
