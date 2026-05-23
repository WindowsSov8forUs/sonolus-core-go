package resource

import (
	"github.com/WindowsSov8forUs/sonolus-core-go/core"
)

const ReplayItemVersion = 1

type ReplayItem struct {
	Name          string     `json:"name"`
	Source        string     `json:"source,omitempty"`
	Version       int        `json:"version"`
	Title         string     `json:"title"`
	Subtitle      string     `json:"subtitle"`
	Author        string     `json:"author"`
	AuthorUser    *UserItem  `json:"authorUser,omitempty"`
	Tags          []core.Tag `json:"tags"`
	Level         LevelItem  `json:"level"`
	Data          core.Srl   `json:"data"`
	Configuration core.Srl   `json:"configuration"`
}

type ReplayConfiguration struct {
	Options     []float64   `json:"options"`
	OptionNames []core.Text `json:"optionNames,omitempty"`
}

type ReplayData struct {
	StartTime   float64             `json:"startTime"`
	SaveTime    float64             `json:"saveTime"`
	Duration    float64             `json:"duration"`
	InputOffset float64             `json:"inputOffset"`
	PlayArea    ReplayDataPlayArea  `json:"playArea"`
	Result      core.GameplayResult `json:"result"`
	Inputs      ReplayDataInputs    `json:"inputs"`
	Entities    []ReplayDataEntity  `json:"entities"`
	Touches     ReplayDataTouches   `json:"touches"`
	Streams     []ReplayDataStream  `json:"streams,omitempty"`
}

type ReplayDataPlayArea struct {
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
}

type ReplayDataInputs struct {
	EntityIndex []float64 `json:"entityIndex"`
	Time        []float64 `json:"time"`
	Judgment    []float64 `json:"judgment"`
	Accuracy    []float64 `json:"accuracy"`
}

type ReplayDataEntity struct {
	Data []ReplayDataEntityData `json:"data"`
}

type ReplayDataEntityData struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

type ReplayDataTouches struct {
	L []float64 `json:"l"`
	T []float64 `json:"t"`
	X []float64 `json:"x"`
	Y []float64 `json:"y"`
}

type ReplayDataStream struct {
	ID     float64   `json:"id"`
	Keys   []float64 `json:"keys"`
	Values []float64 `json:"values"`
}
