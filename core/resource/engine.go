package resource

import (
	"encoding/json"

	"github.com/WindowsSov8forUs/sonolus-core-go/core"

	"github.com/WindowsSov8forUs/sonolus-core-go/errors"
)

type EngineConfigurationOptionType string

type EngineConfigurationOption interface {
	engineConfigurationOption()
	OptionType() EngineConfigurationOptionType
}

const (
	EngineConfigurationOptionTypeSlider EngineConfigurationOptionType = "slider"
	EngineConfigurationOptionTypeToggle EngineConfigurationOptionType = "toggle"
	EngineConfigurationOptionTypeSelect EngineConfigurationOptionType = "select"
)

type EngineConfigurationOptionBase struct {
	Name        core.Text `json:"name"`
	Description string    `json:"description,omitempty"`
	Standard    bool      `json:"standard,omitempty"`
	Advanced    bool      `json:"advanced,omitempty"`
	Scope       string    `json:"scope,omitempty"`
}

func (EngineConfigurationOptionBase) engineConfigurationOption() {}

type EngineConfigurationSliderOption struct {
	EngineConfigurationOptionBase

	Type EngineConfigurationOptionType `json:"type"`
	Def  float64                       `json:"def"`
	Min  float64                       `json:"min"`
	Max  float64                       `json:"max"`
	Step float64                       `json:"step"`
	Unit core.Text                     `json:"unit,omitempty"`
}

func (EngineConfigurationSliderOption) OptionType() EngineConfigurationOptionType {
	return EngineConfigurationOptionTypeSlider
}

type EngineConfigurationToggleOption struct {
	EngineConfigurationOptionBase

	Type EngineConfigurationOptionType `json:"type"`
	Def  int                           `json:"def"`
}

func (EngineConfigurationToggleOption) OptionType() EngineConfigurationOptionType {
	return EngineConfigurationOptionTypeToggle
}

type EngineConfigurationSelectOption struct {
	EngineConfigurationOptionBase

	Type   EngineConfigurationOptionType `json:"type"`
	Def    int                           `json:"def"`
	Values []core.Text                   `json:"values"`
}

func (EngineConfigurationSelectOption) OptionType() EngineConfigurationOptionType {
	return EngineConfigurationOptionTypeSelect
}

type EngineConfigurationMetric string

const (
	EngineConfigurationMetricArcade                  EngineConfigurationMetric = "arcade"
	EngineConfigurationMetricArcadePercentage        EngineConfigurationMetric = "arcadePercentage"
	EngineConfigurationMetricAccuracy                EngineConfigurationMetric = "accuracy"
	EngineConfigurationMetricAccuracyPercentage      EngineConfigurationMetric = "accuracyPercentage"
	EngineConfigurationMetricLife                    EngineConfigurationMetric = "life"
	EngineConfigurationMetricPerfect                 EngineConfigurationMetric = "perfect"
	EngineConfigurationMetricPerfectPercentage       EngineConfigurationMetric = "perfectPercentage"
	EngineConfigurationMetricGreatGoodMiss           EngineConfigurationMetric = "greatGoodMiss"
	EngineConfigurationMetricGreatGoodMissPercentage EngineConfigurationMetric = "greatGoodMissPercentage"
	EngineConfigurationMetricMiss                    EngineConfigurationMetric = "miss"
	EngineConfigurationMetricMissPercentage          EngineConfigurationMetric = "missPercentage"
	EngineConfigurationMetricErrorHeatmap            EngineConfigurationMetric = "errorHeatmap"
)

func (m EngineConfigurationMetric) valid() bool {
	switch m {
	case EngineConfigurationMetricArcade, EngineConfigurationMetricArcadePercentage:
		return true
	case EngineConfigurationMetricAccuracy, EngineConfigurationMetricAccuracyPercentage:
		return true
	case EngineConfigurationMetricLife:
		return true
	case EngineConfigurationMetricPerfect, EngineConfigurationMetricPerfectPercentage:
		return true
	case EngineConfigurationMetricGreatGoodMiss, EngineConfigurationMetricGreatGoodMissPercentage:
		return true
	case EngineConfigurationMetricMiss, EngineConfigurationMetricMissPercentage:
		return true
	case EngineConfigurationMetricErrorHeatmap:
		return true
	default:
		return false
	}
}

func (m *EngineConfigurationMetric) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	value := EngineConfigurationMetric(s)
	if !value.valid() {
		return errors.InvalidEnumValueError{
			TypeName: "EngineConfigurationMetric",
			Value:    string(value),
		}
	}
	*m = value
	return nil
}

type EngineConfigurationVisibility struct {
	Scale float64 `json:"scale"`
	Alpha float64 `json:"alpha"`
}

type EngineConfigurationAnimationTweenEase string

const (
	EngineConfigurationAnimationTweenEaseLinear EngineConfigurationAnimationTweenEase = "linear"

	EngineConfigurationAnimationTweenEaseInSine    EngineConfigurationAnimationTweenEase = "inSine"
	EngineConfigurationAnimationTweenEaseOutSine   EngineConfigurationAnimationTweenEase = "outSine"
	EngineConfigurationAnimationTweenEaseInOutSine EngineConfigurationAnimationTweenEase = "inOutSine"
	EngineConfigurationAnimationTweenEaseOutInSine EngineConfigurationAnimationTweenEase = "outInSine"

	EngineConfigurationAnimationTweenEaseInQuad    EngineConfigurationAnimationTweenEase = "inQuad"
	EngineConfigurationAnimationTweenEaseOutQuad   EngineConfigurationAnimationTweenEase = "outQuad"
	EngineConfigurationAnimationTweenEaseInOutQuad EngineConfigurationAnimationTweenEase = "inOutQuad"
	EngineConfigurationAnimationTweenEaseOutInQuad EngineConfigurationAnimationTweenEase = "outInQuad"

	EngineConfigurationAnimationTweenEaseInCubic    EngineConfigurationAnimationTweenEase = "inCubic"
	EngineConfigurationAnimationTweenEaseOutCubic   EngineConfigurationAnimationTweenEase = "outCubic"
	EngineConfigurationAnimationTweenEaseInOutCubic EngineConfigurationAnimationTweenEase = "inOutCubic"
	EngineConfigurationAnimationTweenEaseOutInCubic EngineConfigurationAnimationTweenEase = "outInCubic"

	EngineConfigurationAnimationTweenEaseInQuart    EngineConfigurationAnimationTweenEase = "inQuart"
	EngineConfigurationAnimationTweenEaseOutQuart   EngineConfigurationAnimationTweenEase = "outQuart"
	EngineConfigurationAnimationTweenEaseInOutQuart EngineConfigurationAnimationTweenEase = "inOutQuart"
	EngineConfigurationAnimationTweenEaseOutInQuart EngineConfigurationAnimationTweenEase = "outInQuart"

	EngineConfigurationAnimationTweenEaseInQuint    EngineConfigurationAnimationTweenEase = "inQuint"
	EngineConfigurationAnimationTweenEaseOutQuint   EngineConfigurationAnimationTweenEase = "outQuint"
	EngineConfigurationAnimationTweenEaseInOutQuint EngineConfigurationAnimationTweenEase = "inOutQuint"
	EngineConfigurationAnimationTweenEaseOutInQuint EngineConfigurationAnimationTweenEase = "outInQuint"

	EngineConfigurationAnimationTweenEaseInExpo    EngineConfigurationAnimationTweenEase = "inExpo"
	EngineConfigurationAnimationTweenEaseOutExpo   EngineConfigurationAnimationTweenEase = "outExpo"
	EngineConfigurationAnimationTweenEaseInOutExpo EngineConfigurationAnimationTweenEase = "inOutExpo"
	EngineConfigurationAnimationTweenEaseOutInExpo EngineConfigurationAnimationTweenEase = "outInExpo"

	EngineConfigurationAnimationTweenEaseInCirc    EngineConfigurationAnimationTweenEase = "inCirc"
	EngineConfigurationAnimationTweenEaseOutCirc   EngineConfigurationAnimationTweenEase = "outCirc"
	EngineConfigurationAnimationTweenEaseInOutCirc EngineConfigurationAnimationTweenEase = "inOutCirc"
	EngineConfigurationAnimationTweenEaseOutInCirc EngineConfigurationAnimationTweenEase = "outInCirc"

	EngineConfigurationAnimationTweenEaseInBack    EngineConfigurationAnimationTweenEase = "inBack"
	EngineConfigurationAnimationTweenEaseOutBack   EngineConfigurationAnimationTweenEase = "outBack"
	EngineConfigurationAnimationTweenEaseInOutBack EngineConfigurationAnimationTweenEase = "inOutBack"
	EngineConfigurationAnimationTweenEaseOutInBack EngineConfigurationAnimationTweenEase = "outInBack"

	EngineConfigurationAnimationTweenEaseInElastic    EngineConfigurationAnimationTweenEase = "inElastic"
	EngineConfigurationAnimationTweenEaseOutElastic   EngineConfigurationAnimationTweenEase = "outElastic"
	EngineConfigurationAnimationTweenEaseInOutElastic EngineConfigurationAnimationTweenEase = "inOutElastic"
	EngineConfigurationAnimationTweenEaseOutInElastic EngineConfigurationAnimationTweenEase = "outInElastic"

	EngineConfigurationAnimationTweenEaseNone EngineConfigurationAnimationTweenEase = "none"
)

func (e EngineConfigurationAnimationTweenEase) valid() bool {
	switch e {
	case EngineConfigurationAnimationTweenEaseLinear:
		return true
	case EngineConfigurationAnimationTweenEaseInSine, EngineConfigurationAnimationTweenEaseOutSine:
		return true
	case EngineConfigurationAnimationTweenEaseInOutSine, EngineConfigurationAnimationTweenEaseOutInSine:
		return true
	case EngineConfigurationAnimationTweenEaseInQuad, EngineConfigurationAnimationTweenEaseOutQuad:
		return true
	case EngineConfigurationAnimationTweenEaseInOutQuad, EngineConfigurationAnimationTweenEaseOutInQuad:
		return true
	case EngineConfigurationAnimationTweenEaseInCubic, EngineConfigurationAnimationTweenEaseOutCubic:
		return true
	case EngineConfigurationAnimationTweenEaseInOutCubic, EngineConfigurationAnimationTweenEaseOutInCubic:
		return true
	case EngineConfigurationAnimationTweenEaseInQuart, EngineConfigurationAnimationTweenEaseOutQuart:
		return true
	case EngineConfigurationAnimationTweenEaseInOutQuart, EngineConfigurationAnimationTweenEaseOutInQuart:
		return true
	case EngineConfigurationAnimationTweenEaseInQuint, EngineConfigurationAnimationTweenEaseOutQuint:
		return true
	case EngineConfigurationAnimationTweenEaseInOutQuint, EngineConfigurationAnimationTweenEaseOutInQuint:
		return true
	case EngineConfigurationAnimationTweenEaseInExpo, EngineConfigurationAnimationTweenEaseOutExpo:
		return true
	case EngineConfigurationAnimationTweenEaseInOutExpo, EngineConfigurationAnimationTweenEaseOutInExpo:
		return true
	case EngineConfigurationAnimationTweenEaseInCirc, EngineConfigurationAnimationTweenEaseOutCirc:
		return true
	case EngineConfigurationAnimationTweenEaseInOutCirc, EngineConfigurationAnimationTweenEaseOutInCirc:
		return true
	case EngineConfigurationAnimationTweenEaseInBack, EngineConfigurationAnimationTweenEaseOutBack:
		return true
	case EngineConfigurationAnimationTweenEaseInOutBack, EngineConfigurationAnimationTweenEaseOutInBack:
		return true
	case EngineConfigurationAnimationTweenEaseInElastic, EngineConfigurationAnimationTweenEaseOutElastic:
		return true
	case EngineConfigurationAnimationTweenEaseInOutElastic, EngineConfigurationAnimationTweenEaseOutInElastic:
		return true
	case EngineConfigurationAnimationTweenEaseNone:
		return true
	default:
		return false
	}
}

func (e *EngineConfigurationAnimationTweenEase) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	value := EngineConfigurationAnimationTweenEase(s)
	if !value.valid() {
		return errors.InvalidEnumValueError{
			TypeName: "EngineConfigurationAnimationTweenEase",
			Value:    string(value),
		}
	}
	*e = value
	return nil
}

type EngineConfigurationAnimationTween struct {
	From     float64                               `json:"from"`
	To       float64                               `json:"to"`
	Duration float64                               `json:"duration"`
	Ease     EngineConfigurationAnimationTweenEase `json:"ease"`
}

type EngineConfigurationAnimation struct {
	Scale EngineConfigurationAnimationTween `json:"scale"`
	Alpha EngineConfigurationAnimationTween `json:"alpha"`
}

type EngineConfigurationJudgmentErrorStyle string

const (
	EngineConfigurationJudgmentErrorStyleNone          EngineConfigurationJudgmentErrorStyle = "none"
	EngineConfigurationJudgmentErrorStyleLate          EngineConfigurationJudgmentErrorStyle = "late"
	EngineConfigurationJudgmentErrorStyleEarly         EngineConfigurationJudgmentErrorStyle = "early"
	EngineConfigurationJudgmentErrorStylePlus          EngineConfigurationJudgmentErrorStyle = "plus"
	EngineConfigurationJudgmentErrorStyleMinus         EngineConfigurationJudgmentErrorStyle = "minus"
	EngineConfigurationJudgmentErrorStyleArrowUp       EngineConfigurationJudgmentErrorStyle = "arrowUp"
	EngineConfigurationJudgmentErrorStyleArrowDown     EngineConfigurationJudgmentErrorStyle = "arrowDown"
	EngineConfigurationJudgmentErrorStyleArrowLeft     EngineConfigurationJudgmentErrorStyle = "arrowLeft"
	EngineConfigurationJudgmentErrorStyleArrowRight    EngineConfigurationJudgmentErrorStyle = "arrowRight"
	EngineConfigurationJudgmentErrorStyleTriangleUp    EngineConfigurationJudgmentErrorStyle = "triangleUp"
	EngineConfigurationJudgmentErrorStyleTriangleDown  EngineConfigurationJudgmentErrorStyle = "triangleDown"
	EngineConfigurationJudgmentErrorStyleTriangleLeft  EngineConfigurationJudgmentErrorStyle = "triangleLeft"
	EngineConfigurationJudgmentErrorStyleTriangleRight EngineConfigurationJudgmentErrorStyle = "triangleRight"
)

func (s EngineConfigurationJudgmentErrorStyle) valid() bool {
	switch s {
	case EngineConfigurationJudgmentErrorStyleNone:
		return true
	case EngineConfigurationJudgmentErrorStyleLate, EngineConfigurationJudgmentErrorStyleEarly:
		return true
	case EngineConfigurationJudgmentErrorStylePlus, EngineConfigurationJudgmentErrorStyleMinus:
		return true
	case EngineConfigurationJudgmentErrorStyleArrowUp, EngineConfigurationJudgmentErrorStyleArrowDown:
		return true
	case EngineConfigurationJudgmentErrorStyleArrowLeft, EngineConfigurationJudgmentErrorStyleArrowRight:
		return true
	case EngineConfigurationJudgmentErrorStyleTriangleUp, EngineConfigurationJudgmentErrorStyleTriangleDown:
		return true
	case EngineConfigurationJudgmentErrorStyleTriangleLeft, EngineConfigurationJudgmentErrorStyleTriangleRight:
		return true
	default:
		return false
	}
}

func (s *EngineConfigurationJudgmentErrorStyle) UnmarshalJSON(data []byte) error {
	var value string
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	style := EngineConfigurationJudgmentErrorStyle(value)
	if !style.valid() {
		return errors.InvalidEnumValueError{
			TypeName: "EngineConfigurationJudgmentErrorStyle",
			Value:    string(style),
		}
	}
	*s = style
	return nil
}

type EngineConfigurationJudgmentErrorPlacement string

const (
	EngineConfigurationJudgmentErrorPlacementLeft      EngineConfigurationJudgmentErrorPlacement = "left"
	EngineConfigurationJudgmentErrorPlacementRight     EngineConfigurationJudgmentErrorPlacement = "right"
	EngineConfigurationJudgmentErrorPlacementLeftRight EngineConfigurationJudgmentErrorPlacement = "leftRight"
	EngineConfigurationJudgmentErrorPlacementTop       EngineConfigurationJudgmentErrorPlacement = "top"
	EngineConfigurationJudgmentErrorPlacementBottom    EngineConfigurationJudgmentErrorPlacement = "bottom"
	EngineConfigurationJudgmentErrorPlacementTopBottom EngineConfigurationJudgmentErrorPlacement = "topBottom"
	EngineConfigurationJudgmentErrorPlacementCenter    EngineConfigurationJudgmentErrorPlacement = "center"
)

func (p EngineConfigurationJudgmentErrorPlacement) valid() bool {
	switch p {
	case EngineConfigurationJudgmentErrorPlacementLeft, EngineConfigurationJudgmentErrorPlacementRight:
		return true
	case EngineConfigurationJudgmentErrorPlacementLeftRight:
		return true
	case EngineConfigurationJudgmentErrorPlacementTop, EngineConfigurationJudgmentErrorPlacementBottom:
		return true
	case EngineConfigurationJudgmentErrorPlacementTopBottom:
		return true
	case EngineConfigurationJudgmentErrorPlacementCenter:
		return true
	default:
		return false
	}
}

func (p *EngineConfigurationJudgmentErrorPlacement) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	value := EngineConfigurationJudgmentErrorPlacement(s)
	if !value.valid() {
		return errors.InvalidEnumValueError{
			TypeName: "EngineConfigurationJudgmentErrorPlacement",
			Value:    string(value),
		}
	}
	*p = value
	return nil
}

type EngineConfigurationUI struct {
	Scope                         string                                    `json:"scope,omitempty"`
	PrimaryMetric                 EngineConfigurationMetric                 `json:"primaryMetric"`
	SecondaryMetric               EngineConfigurationMetric                 `json:"secondaryMetric"`
	MenuVisibility                EngineConfigurationVisibility             `json:"menuVisibility"`
	JudgmentVisibility            EngineConfigurationVisibility             `json:"judgmentVisibility"`
	ComboVisibility               EngineConfigurationVisibility             `json:"comboVisibility"`
	PrimaryMetricVisibility       EngineConfigurationVisibility             `json:"primaryMetricVisibility"`
	SecondaryMetricVisibility     EngineConfigurationVisibility             `json:"secondaryMetricVisibility"`
	ProgressVisibility            EngineConfigurationVisibility             `json:"progressVisibility"`
	TutorialNavigationVisibility  EngineConfigurationVisibility             `json:"tutorialNavigationVisibility"`
	TutorialInstructionVisibility EngineConfigurationVisibility             `json:"tutorialInstructionVisibility"`
	JudgmentAnimation             EngineConfigurationAnimation              `json:"judgmentAnimation"`
	ComboAnimation                EngineConfigurationAnimation              `json:"comboAnimation"`
	JudgmentErrorStyle            EngineConfigurationJudgmentErrorStyle     `json:"judgmentErrorStyle"`
	JudgmentErrorPlacement        EngineConfigurationJudgmentErrorPlacement `json:"judgmentErrorPlacement"`
	JudgmentErrorMin              float64                                   `json:"judgmentErrorMin"`
}

type EngineConfiguration struct {
	Options                   []EngineConfigurationOption `json:"options"`
	UI                        EngineConfigurationUI       `json:"ui"`
	ReplayFallbackOptionNames []core.Text                 `json:"replayFallbackOptionNames,omitempty"`
}

const EngineItemVersion = 13

type EngineArchetypeName string

const (
	EngineArchetypeNameBPMChange       EngineArchetypeName = "#BPM_CHANGE"
	EngineArchetypeNameTimeScaleChange EngineArchetypeName = "#TIMESCALE_CHANGE"
)

type EngineArchetypeDataName string

const (
	EngineArchetypeDataNameBeat      EngineArchetypeDataName = "#BEAT"
	EngineArchetypeDataNameBPM       EngineArchetypeDataName = "#BPM"
	EngineArchetypeDataNameTimeScale EngineArchetypeDataName = "#TIMESCALE"
	EngineArchetypeDataNameJudgment  EngineArchetypeDataName = "#JUDGMENT"
	EngineArchetypeDataNameAccuracy  EngineArchetypeDataName = "#ACCURACY"
)

type EngineDataBucket struct {
	Sprites []EngineDataBucketSprite `json:"sprites"`
	Unit    core.Text                `json:"unit,omitempty"`
}

type EngineDataBucketSprite struct {
	ID         int     `json:"id"`
	FallbackID int     `json:"fallbackId,omitempty"`
	X          float64 `json:"x"`
	Y          float64 `json:"y"`
	W          float64 `json:"w"`
	H          float64 `json:"h"`
	Rotation   float64 `json:"rotation"`
}

type EngineDataNode interface {
	engineDataNode()
}

type EngineDataValueNode struct {
	Value float64 `json:"value"`
}

func (EngineDataValueNode) engineDataNode() {}

type EngineDataFunctionNode struct {
	Func RuntimeFunction `json:"func"`
	Args []int           `json:"args"`
}

func (EngineDataFunctionNode) engineDataNode() {}

type EngineRenderMode string

const (
	EngineRenderModeDefault     EngineRenderMode = "default"
	EngineRenderModeStandard    EngineRenderMode = "standard"
	EngineRenderModeLightweight EngineRenderMode = "lightweight"
)

func (m EngineRenderMode) valid() bool {
	switch m {
	case EngineRenderModeDefault, EngineRenderModeStandard, EngineRenderModeLightweight:
		return true
	default:
		return false
	}
}

func (m *EngineRenderMode) UnmarshalJSON(data []byte) error {
	var value string
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	renderMode := EngineRenderMode(value)
	if !renderMode.valid() {
		return errors.InvalidEnumValueError{
			TypeName: "EngineRenderMode",
			Value:    string(renderMode),
		}
	}
	*m = renderMode
	return nil
}

type EngineSkinData struct {
	RenderMode EngineRenderMode       `json:"renderMode,omitempty"`
	Sprites    []EngineSkinDataSprite `json:"sprites"`
}

type EngineSkinDataSprite struct {
	Name SkinSpriteName `json:"name"`
	ID   int            `json:"id"`
}

type EngineEffectData struct {
	Clips []EngineEffectDataClip `json:"clips"`
}

type EngineEffectDataClip struct {
	Name EffectClipName `json:"name"`
	ID   int            `json:"id"`
}

type EngineParticleData struct {
	Effects []EngineParticleDataEffect `json:"effects"`
}

type EngineParticleDataEffect struct {
	Name ParticleEffectName `json:"name"`
	ID   int                `json:"id"`
}

type EnginePlayData struct {
	Skin       EngineSkinData            `json:"skin"`
	Effect     EngineEffectData          `json:"effect"`
	Particle   EngineParticleData        `json:"particle"`
	Buckets    []EngineDataBucket        `json:"buckets"`
	Archetypes []EnginePlayDataArchetype `json:"archetypes"`
	Nodes      []EngineDataNode          `json:"nodes"`
}

type EnginePlayDataArchetype struct {
	Name             EngineArchetypeName              `json:"name"`
	HasInput         bool                             `json:"hasInput"`
	Preprocess       *EnginePlayDataArchetypeCallback `json:"preprocess,omitempty"`
	SpawnOrder       *EnginePlayDataArchetypeCallback `json:"spawnOrder,omitempty"`
	ShouldSpawn      *EnginePlayDataArchetypeCallback `json:"shouldSpawn,omitempty"`
	Initialize       *EnginePlayDataArchetypeCallback `json:"initialize,omitempty"`
	UpdateSequential *EnginePlayDataArchetypeCallback `json:"updateSequential,omitempty"`
	Touch            *EnginePlayDataArchetypeCallback `json:"touch,omitempty"`
	UpdateParallel   *EnginePlayDataArchetypeCallback `json:"updateParallel,omitempty"`
	Terminate        *EnginePlayDataArchetypeCallback `json:"terminate,omitempty"`
	Imports          []EngineDataArchetypeImport      `json:"imports"`
	Exports          []EngineArchetypeDataName        `json:"exports"`
}

type EnginePlayDataArchetypeCallback struct {
	Index int `json:"index"`
	Order int `json:"order,omitempty"`
}

type EngineDataArchetypeImport struct {
	Name  EngineArchetypeDataName `json:"name"`
	Index int                     `json:"index"`
	Def   float64                 `json:"def,omitempty"`
}

type EngineWatchData struct {
	Skin        EngineSkinData             `json:"skin"`
	Effect      EngineEffectData           `json:"effect"`
	Particle    EngineParticleData         `json:"particle"`
	Buckets     []EngineDataBucket         `json:"buckets"`
	Archetypes  []EngineWatchDataArchetype `json:"archetypes"`
	UpdateSpawn int                        `json:"updateSpawn,omitempty"`
	Nodes       []EngineDataNode           `json:"nodes"`
}

type EngineWatchDataArchetype struct {
	Name             EngineArchetypeName               `json:"name"`
	HasInput         bool                              `json:"hasInput"`
	Preprocess       *EngineWatchDataArchetypeCallback `json:"preprocess,omitempty"`
	SpawnTime        *EngineWatchDataArchetypeCallback `json:"spawnTime,omitempty"`
	DespawnTime      *EngineWatchDataArchetypeCallback `json:"despawnTime,omitempty"`
	Initialize       *EngineWatchDataArchetypeCallback `json:"initialize,omitempty"`
	UpdateSequential *EngineWatchDataArchetypeCallback `json:"updateSequential,omitempty"`
	UpdateParallel   *EngineWatchDataArchetypeCallback `json:"updateParallel,omitempty"`
	Terminate        *EngineWatchDataArchetypeCallback `json:"terminate,omitempty"`
	Imports          []EngineDataArchetypeImport       `json:"imports"`
}

type EngineWatchDataArchetypeCallback struct {
	Index int `json:"index"`
	Order int `json:"order,omitempty"`
}

type EnginePreviewData struct {
	Skin       EngineSkinData               `json:"skin"`
	Archetypes []EnginePreviewDataArchetype `json:"archetypes"`
	Nodes      []EngineDataNode             `json:"nodes"`
}

type EnginePreviewDataArchetype struct {
	Name       EngineArchetypeName                 `json:"name"`
	Preprocess *EnginePreviewDataArchetypeCallback `json:"preprocess,omitempty"`
	Render     *EnginePreviewDataArchetypeCallback `json:"render,omitempty"`
	Imports    []EngineDataArchetypeImport         `json:"imports"`
}

type EnginePreviewDataArchetypeCallback struct {
	Index int `json:"index"`
	Order int `json:"order,omitempty"`
}

type EngineTutorialData struct {
	Skin        EngineSkinData        `json:"skin"`
	Effect      EngineEffectData      `json:"effect"`
	Particle    EngineParticleData    `json:"particle"`
	Instruction EngineInstructionData `json:"instruction"`
	Preprocess  int                   `json:"preprocess,omitempty"`
	Navigate    int                   `json:"navigate,omitempty"`
	Update      int                   `json:"update,omitempty"`
	Nodes       []EngineDataNode      `json:"nodes"`
}

type EngineInstructionData struct {
	Texts []EngineInstructionDataText `json:"texts"`
	Icons []EngineInstructionDataIcon `json:"icons"`
}

type EngineInstructionDataText struct {
	Name core.Text `json:"name"`
	ID   int       `json:"id"`
}

type EngineInstructionDataIcon struct {
	Name InstructionIconName `json:"name"`
	ID   int                 `json:"id"`
}

type EngineItem struct {
	Name          string         `json:"name"`
	Source        string         `json:"source,omitempty"`
	Version       int            `json:"version"`
	Title         string         `json:"title"`
	Subtitle      string         `json:"subtitle"`
	Author        string         `json:"author"`
	AuthorUser    *UserItem      `json:"authorUser,omitempty"`
	Tags          []core.Tag     `json:"tags"`
	Skin          SkinItem       `json:"skin"`
	Background    BackgroundItem `json:"background"`
	Effect        EffectItem     `json:"effect"`
	Particle      ParticleItem   `json:"particle"`
	Thumbnail     core.Srl       `json:"thumbnail"`
	PlayData      core.Srl       `json:"playData"`
	WatchData     core.Srl       `json:"watchData"`
	PreviewData   core.Srl       `json:"previewData"`
	TutorialData  core.Srl       `json:"tutorialData"`
	ROM           *core.Srl      `json:"rom,omitempty"`
	Configuration core.Srl       `json:"configuration"`
}
