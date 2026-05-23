package database

import "github.com/WindowsSov8forUs/sonolus-core-go/core"

const DatabaseBackgroundItemVersion = 2

type DatabaseBackgroundItem struct {
	Name          string           `json:"name"`
	Version       int              `json:"version"`
	Title         LocalizationText `json:"title"`
	Subtitle      LocalizationText `json:"subtitle"`
	Author        LocalizationText `json:"author"`
	Tags          []DatabaseTag    `json:"tags"`
	Description   LocalizationText `json:"description,omitempty"`
	Thumbnail     core.Srl         `json:"thumbnail"`
	Data          core.Srl         `json:"data"`
	Image         core.Srl         `json:"image"`
	Configuration core.Srl         `json:"configuration"`
}
