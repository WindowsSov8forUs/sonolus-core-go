package database

import "github.com/WindowsSov8forUs/sonolus-core-go/core"

const DatabaseEngineItemVersion = 13

type DatabaseEngineItem struct {
	Name          string           `json:"name"`
	Version       int              `json:"version"`
	Title         LocalizationText `json:"title"`
	Subtitle      LocalizationText `json:"subtitle"`
	Author        LocalizationText `json:"author"`
	Tags          []DatabaseTag    `json:"tags"`
	Description   LocalizationText `json:"description,omitempty"`
	Skin          string           `json:"skin"`
	Background    string           `json:"background"`
	Effect        string           `json:"effect"`
	Particle      string           `json:"particle"`
	Thumbnail     core.Srl         `json:"thumbnail"`
	PlayData      core.Srl         `json:"playData"`
	WatchData     core.Srl         `json:"watchData"`
	PreviewData   core.Srl         `json:"previewData"`
	TutorialData  core.Srl         `json:"tutorialData"`
	ROM           *core.Srl        `json:"rom,omitempty"`
	Configuration core.Srl         `json:"configuration"`
}
