package core

import (
	"encoding/json"

	"github.com/WindowsSov8forUs/sonolus-core-go/errors"
)

type AutoExit string

const (
	AutoExitOff        AutoExit = "off"
	AutoExitPass       AutoExit = "pass"
	AutoExitFullCombo  AutoExit = "fullCombo"
	AutoExitAllPerfect AutoExit = "allPerfect"
)

func (ae AutoExit) valid() bool {
	switch ae {
	case AutoExitOff, AutoExitPass, AutoExitFullCombo, AutoExitAllPerfect:
		return true
	default:
		return false
	}
}

func (ae *AutoExit) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	value := AutoExit(s)
	if !value.valid() {
		return errors.InvalidEnumValueError{
			TypeName: "AutoExit",
			Value:    string(value),
		}
	}
	*ae = value
	return nil
}

type Grade string

const (
	GradeAllPerfect Grade = "allPerfect"
	GradeFullCombo  Grade = "fullCombo"
	GradePass       Grade = "pass"
	GradeFail       Grade = "fail"
)

func (g Grade) valid() bool {
	switch g {
	case GradeAllPerfect, GradeFullCombo, GradePass, GradeFail:
		return true
	default:
		return false
	}
}

func (g *Grade) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	value := Grade(s)
	if !value.valid() {
		return errors.InvalidEnumValueError{
			TypeName: "Grade",
			Value:    string(value),
		}
	}
	*g = value
	return nil
}

type ItemType string

const (
	ItemTypePost       ItemType = "post"
	ItemTypePlaylist   ItemType = "playlist"
	ItemTypeLevel      ItemType = "level"
	ItemTypeSkin       ItemType = "skin"
	ItemTypeBackground ItemType = "background"
	ItemTypeEffect     ItemType = "effect"
	ItemTypeParticle   ItemType = "particle"
	ItemTypeEngine     ItemType = "engine"
	ItemTypeReplay     ItemType = "replay"
	ItemTypeRoom       ItemType = "room"
	ItemTypeUser       ItemType = "user"
)

func (it ItemType) valid() bool {
	switch it {
	case ItemTypePost, ItemTypePlaylist, ItemTypeLevel:
		return true
	case ItemTypeSkin, ItemTypeBackground, ItemTypeEffect:
		return true
	case ItemTypeParticle, ItemTypeEngine, ItemTypeReplay:
		return true
	case ItemTypeRoom, ItemTypeUser:
		return true
	default:
		return false
	}
}

func (it *ItemType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	value := ItemType(s)
	if !value.valid() {
		return errors.InvalidEnumValueError{
			TypeName: "ItemType",
			Value:    string(value),
		}
	}
	*it = value
	return nil
}
