package database

import "github.com/WindowsSov8forUs/sonolus-core-go/core"

const DatabasePlaylistItemVersion = 1

type DatabasePlaylistItem struct {
	Name        string           `json:"name"`
	Version     int              `json:"version"`
	Title       LocalizationText `json:"title"`
	Subtitle    LocalizationText `json:"subtitle"`
	Author      LocalizationText `json:"author"`
	Tags        []DatabaseTag    `json:"tags"`
	Description LocalizationText `json:"description,omitempty"`
	Levels      []string         `json:"levels"`
	Thumbnail   *core.Srl        `json:"thumbnail,omitempty"`
}
