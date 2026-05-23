package resource

import (
	"github.com/WindowsSov8forUs/sonolus-core-go/core"
)

const ParticleItemVersion = 3

type ParticleEffectName string

const (
	ParticleEffectNameNoteCircularTapNeutral         ParticleEffectName = "#NOTE_CIRCULAR_TAP_NEUTRAL"
	ParticleEffectNameNoteCircularTapRed             ParticleEffectName = "#NOTE_CIRCULAR_TAP_RED"
	ParticleEffectNameNoteCircularTapGreen           ParticleEffectName = "#NOTE_CIRCULAR_TAP_GREEN"
	ParticleEffectNameNoteCircularTapBlue            ParticleEffectName = "#NOTE_CIRCULAR_TAP_BLUE"
	ParticleEffectNameNoteCircularTapYellow          ParticleEffectName = "#NOTE_CIRCULAR_TAP_YELLOW"
	ParticleEffectNameNoteCircularTapPurple          ParticleEffectName = "#NOTE_CIRCULAR_TAP_PURPLE"
	ParticleEffectNameNoteCircularTapCyan            ParticleEffectName = "#NOTE_CIRCULAR_TAP_CYAN"
	ParticleEffectNameNoteCircularAlternativeNeutral ParticleEffectName = "#NOTE_CIRCULAR_ALTERNATIVE_NEUTRAL"
	ParticleEffectNameNoteCircularAlternativeRed     ParticleEffectName = "#NOTE_CIRCULAR_ALTERNATIVE_RED"
	ParticleEffectNameNoteCircularAlternativeGreen   ParticleEffectName = "#NOTE_CIRCULAR_ALTERNATIVE_GREEN"
	ParticleEffectNameNoteCircularAlternativeBlue    ParticleEffectName = "#NOTE_CIRCULAR_ALTERNATIVE_BLUE"
	ParticleEffectNameNoteCircularAlternativeYellow  ParticleEffectName = "#NOTE_CIRCULAR_ALTERNATIVE_YELLOW"
	ParticleEffectNameNoteCircularAlternativePurple  ParticleEffectName = "#NOTE_CIRCULAR_ALTERNATIVE_PURPLE"
	ParticleEffectNameNoteCircularAlternativeCyan    ParticleEffectName = "#NOTE_CIRCULAR_ALTERNATIVE_CYAN"
	ParticleEffectNameNoteCircularHoldNeutral        ParticleEffectName = "#NOTE_CIRCULAR_HOLD_NEUTRAL"
	ParticleEffectNameNoteCircularHoldRed            ParticleEffectName = "#NOTE_CIRCULAR_HOLD_RED"
	ParticleEffectNameNoteCircularHoldGreen          ParticleEffectName = "#NOTE_CIRCULAR_HOLD_GREEN"
	ParticleEffectNameNoteCircularHoldBlue           ParticleEffectName = "#NOTE_CIRCULAR_HOLD_BLUE"
	ParticleEffectNameNoteCircularHoldYellow         ParticleEffectName = "#NOTE_CIRCULAR_HOLD_YELLOW"
	ParticleEffectNameNoteCircularHoldPurple         ParticleEffectName = "#NOTE_CIRCULAR_HOLD_PURPLE"
	ParticleEffectNameNoteCircularHoldCyan           ParticleEffectName = "#NOTE_CIRCULAR_HOLD_CYAN"
	ParticleEffectNameNoteLinearTapNeutral           ParticleEffectName = "#NOTE_LINEAR_TAP_NEUTRAL"
	ParticleEffectNameNoteLinearTapRed               ParticleEffectName = "#NOTE_LINEAR_TAP_RED"
	ParticleEffectNameNoteLinearTapGreen             ParticleEffectName = "#NOTE_LINEAR_TAP_GREEN"
	ParticleEffectNameNoteLinearTapBlue              ParticleEffectName = "#NOTE_LINEAR_TAP_BLUE"
	ParticleEffectNameNoteLinearTapYellow            ParticleEffectName = "#NOTE_LINEAR_TAP_YELLOW"
	ParticleEffectNameNoteLinearTapPurple            ParticleEffectName = "#NOTE_LINEAR_TAP_PURPLE"
	ParticleEffectNameNoteLinearTapCyan              ParticleEffectName = "#NOTE_LINEAR_TAP_CYAN"
	ParticleEffectNameNoteLinearAlternativeNeutral   ParticleEffectName = "#NOTE_LINEAR_ALTERNATIVE_NEUTRAL"
	ParticleEffectNameNoteLinearAlternativeRed       ParticleEffectName = "#NOTE_LINEAR_ALTERNATIVE_RED"
	ParticleEffectNameNoteLinearAlternativeGreen     ParticleEffectName = "#NOTE_LINEAR_ALTERNATIVE_GREEN"
	ParticleEffectNameNoteLinearAlternativeBlue      ParticleEffectName = "#NOTE_LINEAR_ALTERNATIVE_BLUE"
	ParticleEffectNameNoteLinearAlternativeYellow    ParticleEffectName = "#NOTE_LINEAR_ALTERNATIVE_YELLOW"
	ParticleEffectNameNoteLinearAlternativePurple    ParticleEffectName = "#NOTE_LINEAR_ALTERNATIVE_PURPLE"
	ParticleEffectNameNoteLinearAlternativeCyan      ParticleEffectName = "#NOTE_LINEAR_ALTERNATIVE_CYAN"
	ParticleEffectNameNoteLinearHoldNeutral          ParticleEffectName = "#NOTE_LINEAR_HOLD_NEUTRAL"
	ParticleEffectNameNoteLinearHoldRed              ParticleEffectName = "#NOTE_LINEAR_HOLD_RED"
	ParticleEffectNameNoteLinearHoldGreen            ParticleEffectName = "#NOTE_LINEAR_HOLD_GREEN"
	ParticleEffectNameNoteLinearHoldBlue             ParticleEffectName = "#NOTE_LINEAR_HOLD_BLUE"
	ParticleEffectNameNoteLinearHoldYellow           ParticleEffectName = "#NOTE_LINEAR_HOLD_YELLOW"
	ParticleEffectNameNoteLinearHoldPurple           ParticleEffectName = "#NOTE_LINEAR_HOLD_PURPLE"
	ParticleEffectNameNoteLinearHoldCyan             ParticleEffectName = "#NOTE_LINEAR_HOLD_CYAN"
	ParticleEffectNameLaneCircular                   ParticleEffectName = "#LANE_CIRCULAR"
	ParticleEffectNameLaneLinear                     ParticleEffectName = "#LANE_LINEAR"
	ParticleEffectNameSlotCircular                   ParticleEffectName = "#SLOT_CIRCULAR"
	ParticleEffectNameSlotLinear                     ParticleEffectName = "#SLOT_LINEAR"
	ParticleEffectNameJudgeLineCircular              ParticleEffectName = "#JUDGE_LINE_CIRCULAR"
	ParticleEffectNameJudgeLineLinear                ParticleEffectName = "#JUDGE_LINE_LINEAR"
)

type ParticleItem struct {
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
	Texture    core.Srl   `json:"texture"`
}

type ParticleData struct {
	Width         float64              `json:"width"`
	Height        float64              `json:"height"`
	Interpolation bool                 `json:"interpolation"`
	Sprites       []ParticleDataSprite `json:"sprites"`
	Effects       []ParticleDataEffect `json:"effects"`
}

type ParticleDataSprite struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	W float64 `json:"w"`
	H float64 `json:"h"`
}

type ParticleDataEffect struct {
	Name      ParticleEffectName    `json:"name"`
	Transform ParticleDataTransform `json:"transform"`
	Groups    []ParticleDataGroup   `json:"groups"`
}

type ParticleDataTransform map[string]ParticleDataTransformEntry

type ParticleDataTransformEntry map[string]float64

type ParticleDataGroup struct {
	Count     int                         `json:"count"`
	Particles []ParticleDataGroupParticle `json:"particles"`
}

type ParticleDataGroupParticle struct {
	Sprite   int                               `json:"sprite"`
	Color    string                            `json:"color"`
	Start    float64                           `json:"start"`
	Duration float64                           `json:"duration"`
	X        ParticleDataGroupParticleProperty `json:"x"`
	Y        ParticleDataGroupParticleProperty `json:"y"`
	W        ParticleDataGroupParticleProperty `json:"w"`
	H        ParticleDataGroupParticleProperty `json:"h"`
	R        ParticleDataGroupParticleProperty `json:"r"`
	A        ParticleDataGroupParticleProperty `json:"a"`
}

type ParticleDataGroupParticleProperty struct {
	From *ParticleDataGroupParticlePropertyExpression `json:"from,omitempty"`
	To   *ParticleDataGroupParticlePropertyExpression `json:"to,omitempty"`
	Ease EngineConfigurationAnimationTweenEase        `json:"ease,omitempty"`
}

type ParticleDataGroupParticlePropertyExpression map[string]float64
