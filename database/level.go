package database

import "github.com/WindowsSov8forUs/sonolus-core-go/core"

const DatabaseLevelItemVersion = 1

type DatabaseUseItem struct {
	UseDefault bool   `json:"useDefault"`
	Item       string `json:"item,omitempty"`
}

type DatabaseLevelItem struct {
	Name          string           `json:"name"`
	Version       int              `json:"version"`
	Rating        float64          `json:"rating"`
	Title         LocalizationText `json:"title"`
	Artists       LocalizationText `json:"artists"`
	Author        LocalizationText `json:"author"`
	Tags          []DatabaseTag    `json:"tags"`
	Description   LocalizationText `json:"description,omitempty"`
	Engine        string           `json:"engine"`
	UseSkin       DatabaseUseItem  `json:"useSkin"`
	UseBackground DatabaseUseItem  `json:"useBackground"`
	UseEffect     DatabaseUseItem  `json:"useEffect"`
	UseParticle   DatabaseUseItem  `json:"useParticle"`
	Cover         core.Srl         `json:"cover"`
	BGM           core.Srl         `json:"bgm"`
	Preview       *core.Srl        `json:"preview,omitempty"`
	Data          core.Srl         `json:"data"`
}
