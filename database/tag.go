package database

import "github.com/WindowsSov8forUs/sonolus-core-go/core"

type DatabaseTag struct {
	Title LocalizationText `json:"title,omitempty"`
	Icon  core.Icon        `json:"icon,omitempty"`
}
