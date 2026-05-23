package content

import (
	"encoding/json"
	"github.com/WindowsSov8forUs/sonolus-core-go/core"

	"github.com/WindowsSov8forUs/sonolus-core-go/errors"
)

type Fit string

const (
	FitWidth   Fit = "width"
	FitHeight  Fit = "height"
	FitContain Fit = "contain"
	FitCover   Fit = "cover"
)

func (f Fit) valid() bool {
	switch f {
	case FitWidth, FitHeight, FitContain, FitCover:
		return true
	default:
		return false
	}
}

func (f *Fit) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	value := Fit(s)
	if !value.valid() {
		return errors.InvalidEnumValueError{
			TypeName: "Fit",
			Value:    string(value),
		}
	}
	*f = value
	return nil
}

const BackgroundItemVersion = 2

type BackgroundItem struct {
	Name          string     `json:"name"`
	Source        string     `json:"source,omitempty"`
	Version       int        `json:"version"`
	Title         string     `json:"title"`
	Subtitle      string     `json:"subtitle"`
	Author        string     `json:"author"`
	AuthorUser    *UserItem  `json:"authorUser,omitempty"`
	Tags          []core.Tag `json:"tags"`
	Thumbnail     core.Srl   `json:"thumbnail"`
	Data          core.Srl   `json:"data"`
	Image         core.Srl   `json:"image"`
	Configuration core.Srl   `json:"configuration"`
}

type BackgroundConfiguration struct {
	Scope string  `json:"scope,omitempty"`
	Blur  float64 `json:"blur"`
	Mask  string  `json:"mask"`
}

type BackgroundData struct {
	AspectRatio float64 `json:"aspectRatio,omitempty"`
	Fit         Fit     `json:"fit"`
	Color       string  `json:"color"`
}
