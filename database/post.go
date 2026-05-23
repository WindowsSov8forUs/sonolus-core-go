package database

import "github.com/WindowsSov8forUs/sonolus-core-go/core"

const DatabasePostItemVersion = 1

type DatabasePostItem struct {
	Name        string           `json:"name"`
	Version     int              `json:"version"`
	Title       LocalizationText `json:"title"`
	Time        float64          `json:"time"`
	Author      LocalizationText `json:"author"`
	Tags        []DatabaseTag    `json:"tags"`
	Description LocalizationText `json:"description,omitempty"`
	Thumbnail   *core.Srl        `json:"thumbnail,omitempty"`
}
