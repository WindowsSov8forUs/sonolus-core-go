package database

import "github.com/WindowsSov8forUs/sonolus-core-go/core"

const DatabaseSkinItemVersion = 4

type DatabaseSkinItem struct {
	Name        string           `json:"name"`
	Version     int              `json:"version"`
	Title       LocalizationText `json:"title"`
	Subtitle    LocalizationText `json:"subtitle"`
	Author      LocalizationText `json:"author"`
	Tags        []DatabaseTag    `json:"tags"`
	Description LocalizationText `json:"description,omitempty"`
	Thumbnail   core.Srl         `json:"thumbnail"`
	Data        core.Srl         `json:"data"`
	Texture     core.Srl         `json:"texture"`
}
