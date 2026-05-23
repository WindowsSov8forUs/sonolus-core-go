package resource

import (
	"github.com/WindowsSov8forUs/sonolus-core-go/core"
)

const EffectItemVersion = 5

type EffectClipName string

const (
	EffectClipNameMiss    EffectClipName = "#MISS"
	EffectClipNamePerfect EffectClipName = "#PERFECT"
	EffectClipNameGreat   EffectClipName = "#GREAT"
	EffectClipNameGood    EffectClipName = "#GOOD"

	EffectClipNameHold EffectClipName = "#HOLD"

	EffectClipNameMissAlternative    EffectClipName = "#MISS_ALTERNATIVE"
	EffectClipNamePerfectAlternative EffectClipName = "#PERFECT_ALTERNATIVE"
	EffectClipNameGreatAlternative   EffectClipName = "#GREAT_ALTERNATIVE"
	EffectClipNameGoodAlternative    EffectClipName = "#GOOD_ALTERNATIVE"

	EffectClipNameHoldAlternative EffectClipName = "#HOLD_ALTERNATIVE"

	EffectClipNameStage EffectClipName = "#STAGE"
)

type EffectDataClip struct {
	Name     EffectClipName `json:"name"`
	Filename string         `json:"filename"`
}

type EffectData struct {
	Clips []EffectDataClip `json:"clips"`
}

type EffectItem struct {
	Name       string     `json:"name"`
	Source     string     `json:"source,omitempty"`
	Version    int        `json:"version"`
	Title      string     `json:"title"`
	Subtitle   string     `json:"subtitle"`
	Author     string     `json:"author"`
	AuthorUser *UserItem  `json:"authorUser,omitempty"`
	Tags       []core.Tag `json:"tags"`
	Thumbnail  core.Srl   `json:"thumbnail"`
	Data       core.Srl   `json:"data"`
	Audio      core.Srl   `json:"audio"`
}

func DecodeEffectItem(data []byte) (EffectItem, error) {
	return decodeItemAs[EffectItem](data)
}
