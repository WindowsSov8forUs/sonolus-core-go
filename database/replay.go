package database

import "github.com/WindowsSov8forUs/sonolus-core-go/core"

const DatabaseReplayItemVersion = 1

type DatabaseReplayItem struct {
	Name          string           `json:"name"`
	Version       int              `json:"version"`
	Title         LocalizationText `json:"title"`
	Subtitle      LocalizationText `json:"subtitle"`
	Author        LocalizationText `json:"author"`
	Tags          []DatabaseTag    `json:"tags"`
	Description   LocalizationText `json:"description,omitempty"`
	Level         string           `json:"level"`
	Data          core.Srl         `json:"data"`
	Configuration core.Srl         `json:"configuration"`
}
