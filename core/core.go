package core

import "github.com/WindowsSov8forUs/sonolus-core-go/internal/nilable"

type GameplayResult struct {
	Grade         Grade   `json:"grade"`
	ArcadeScore   float64 `json:"arcadeScore"`
	AccuracyScore float64 `json:"accuracyScore"`
	Combo         int     `json:"combo"`
	Perfect       int     `json:"perfect"`
	Great         int     `json:"great"`
	Good          int     `json:"good"`
	Miss          int     `json:"miss"`
	TotalCount    int     `json:"totalCount"`
}

type Sil struct {
	Address string `json:"address"`
	Name    string `json:"name"`
}

type Srl struct {
	Hash *nilable.Nilable[string] `json:"hash,omitempty"`
	URL  *nilable.Nilable[string] `json:"url,omitempty"`
}

type Tag struct {
	Title Text `json:"title,omitempty"`
	Icon  Icon `json:"icon,omitempty"`
}
