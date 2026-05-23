package content

import (
	"github.com/WindowsSov8forUs/sonolus-core-go/core"
)

const LevelItemVersion = 1

type UseItem[T any] struct {
	UseDefault bool `json:"useDefault"`
	Item       *T   `json:"item,omitempty"`
}

type LevelItem struct {
	Name          string                  `json:"name"`
	Source        string                  `json:"source,omitempty"`
	Version       int                     `json:"version"`
	Rating        float64                 `json:"rating"`
	Title         string                  `json:"title"`
	Artists       string                  `json:"artists"`
	Author        string                  `json:"author"`
	AuthorUser    *UserItem               `json:"authorUser,omitempty"`
	Tags          []core.Tag              `json:"tags"`
	Engine        EngineItem              `json:"engine"`
	UseSkin       UseItem[SkinItem]       `json:"useSkin"`
	UseBackground UseItem[BackgroundItem] `json:"useBackground"`
	UseEffect     UseItem[EffectItem]     `json:"useEffect"`
	UseParticle   UseItem[ParticleItem]   `json:"useParticle"`
	Cover         core.Srl                `json:"cover"`
	BGM           core.Srl                `json:"bgm"`
	Preview       *core.Srl               `json:"preview,omitempty"`
	Data          core.Srl                `json:"data"`
}

type LevelData struct {
	BGMOffset float64           `json:"bgmOffset"`
	Entities  []LevelDataEntity `json:"entities"`
}

type LevelDataEntity struct {
	Name      string                `json:"name,omitempty"`
	Archetype EngineArchetypeName   `json:"archetype"`
	Data      []LevelDataEntityData `json:"data"`
}

type LevelDataEntityData interface {
	levelDataEntityData()
}

type LevelDataEntityValueData struct {
	Name  EngineArchetypeDataName `json:"name"`
	Value float64                 `json:"value"`
}

func (LevelDataEntityValueData) levelDataEntityData() {}

type LevelDataEntityRefData struct {
	Name EngineArchetypeDataName `json:"name"`
	Ref  string                  `json:"ref"`
}

func (LevelDataEntityRefData) levelDataEntityData() {}
