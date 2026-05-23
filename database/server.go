package database

import "github.com/WindowsSov8forUs/sonolus-core-go/core"

type DatabaseServerInfo struct {
	Title       LocalizationText `json:"title"`
	Description LocalizationText `json:"description,omitempty"`
	Banner      *core.Srl        `json:"banner,omitempty"`
}
