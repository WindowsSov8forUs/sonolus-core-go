package content

import (
	"github.com/WindowsSov8forUs/sonolus-core-go/core"
)

const SkinItemVersion = 4

type SkinSpriteName string

const (
	SkinSpriteNameNoteHeadNeutral                       SkinSpriteName = "#NOTE_HEAD_NEUTRAL"
	SkinSpriteNameNoteHeadRed                           SkinSpriteName = "#NOTE_HEAD_RED"
	SkinSpriteNameNoteHeadGreen                         SkinSpriteName = "#NOTE_HEAD_GREEN"
	SkinSpriteNameNoteHeadBlue                          SkinSpriteName = "#NOTE_HEAD_BLUE"
	SkinSpriteNameNoteHeadYellow                        SkinSpriteName = "#NOTE_HEAD_YELLOW"
	SkinSpriteNameNoteHeadPurple                        SkinSpriteName = "#NOTE_HEAD_PURPLE"
	SkinSpriteNameNoteHeadCyan                          SkinSpriteName = "#NOTE_HEAD_CYAN"
	SkinSpriteNameNoteTickNeutral                       SkinSpriteName = "#NOTE_TICK_NEUTRAL"
	SkinSpriteNameNoteTickRed                           SkinSpriteName = "#NOTE_TICK_RED"
	SkinSpriteNameNoteTickGreen                         SkinSpriteName = "#NOTE_TICK_GREEN"
	SkinSpriteNameNoteTickBlue                          SkinSpriteName = "#NOTE_TICK_BLUE"
	SkinSpriteNameNoteTickYellow                        SkinSpriteName = "#NOTE_TICK_YELLOW"
	SkinSpriteNameNoteTickPurple                        SkinSpriteName = "#NOTE_TICK_PURPLE"
	SkinSpriteNameNoteTickCyan                          SkinSpriteName = "#NOTE_TICK_CYAN"
	SkinSpriteNameNoteTailNeutral                       SkinSpriteName = "#NOTE_TAIL_NEUTRAL"
	SkinSpriteNameNoteTailRed                           SkinSpriteName = "#NOTE_TAIL_RED"
	SkinSpriteNameNoteTailGreen                         SkinSpriteName = "#NOTE_TAIL_GREEN"
	SkinSpriteNameNoteTailBlue                          SkinSpriteName = "#NOTE_TAIL_BLUE"
	SkinSpriteNameNoteTailYellow                        SkinSpriteName = "#NOTE_TAIL_YELLOW"
	SkinSpriteNameNoteTailPurple                        SkinSpriteName = "#NOTE_TAIL_PURPLE"
	SkinSpriteNameNoteTailCyan                          SkinSpriteName = "#NOTE_TAIL_CYAN"
	SkinSpriteNameNoteConnectionNeutral                 SkinSpriteName = "#NOTE_CONNECTION_NEUTRAL"
	SkinSpriteNameNoteConnectionRed                     SkinSpriteName = "#NOTE_CONNECTION_RED"
	SkinSpriteNameNoteConnectionGreen                   SkinSpriteName = "#NOTE_CONNECTION_GREEN"
	SkinSpriteNameNoteConnectionBlue                    SkinSpriteName = "#NOTE_CONNECTION_BLUE"
	SkinSpriteNameNoteConnectionYellow                  SkinSpriteName = "#NOTE_CONNECTION_YELLOW"
	SkinSpriteNameNoteConnectionPurple                  SkinSpriteName = "#NOTE_CONNECTION_PURPLE"
	SkinSpriteNameNoteConnectionCyan                    SkinSpriteName = "#NOTE_CONNECTION_CYAN"
	SkinSpriteNameNoteConnectionNeutralSeamless         SkinSpriteName = "#NOTE_CONNECTION_NEUTRAL_SEAMLESS"
	SkinSpriteNameNoteConnectionRedSeamless             SkinSpriteName = "#NOTE_CONNECTION_RED_SEAMLESS"
	SkinSpriteNameNoteConnectionGreenSeamless           SkinSpriteName = "#NOTE_CONNECTION_GREEN_SEAMLESS"
	SkinSpriteNameNoteConnectionBlueSeamless            SkinSpriteName = "#NOTE_CONNECTION_BLUE_SEAMLESS"
	SkinSpriteNameNoteConnectionYellowSeamless          SkinSpriteName = "#NOTE_CONNECTION_YELLOW_SEAMLESS"
	SkinSpriteNameNoteConnectionPurpleSeamless          SkinSpriteName = "#NOTE_CONNECTION_PURPLE_SEAMLESS"
	SkinSpriteNameNoteConnectionCyanSeamless            SkinSpriteName = "#NOTE_CONNECTION_CYAN_SEAMLESS"
	SkinSpriteNameSimultaneousConnectionNeutral         SkinSpriteName = "#SIMULTANEOUS_CONNECTION_NEUTRAL"
	SkinSpriteNameSimultaneousConnectionRed             SkinSpriteName = "#SIMULTANEOUS_CONNECTION_RED"
	SkinSpriteNameSimultaneousConnectionGreen           SkinSpriteName = "#SIMULTANEOUS_CONNECTION_GREEN"
	SkinSpriteNameSimultaneousConnectionBlue            SkinSpriteName = "#SIMULTANEOUS_CONNECTION_BLUE"
	SkinSpriteNameSimultaneousConnectionYellow          SkinSpriteName = "#SIMULTANEOUS_CONNECTION_YELLOW"
	SkinSpriteNameSimultaneousConnectionPurple          SkinSpriteName = "#SIMULTANEOUS_CONNECTION_PURPLE"
	SkinSpriteNameSimultaneousConnectionCyan            SkinSpriteName = "#SIMULTANEOUS_CONNECTION_CYAN"
	SkinSpriteNameSimultaneousConnectionNeutralSeamless SkinSpriteName = "#SIMULTANEOUS_CONNECTION_NEUTRAL_SEAMLESS"
	SkinSpriteNameSimultaneousConnectionRedSeamless     SkinSpriteName = "#SIMULTANEOUS_CONNECTION_RED_SEAMLESS"
	SkinSpriteNameSimultaneousConnectionGreenSeamless   SkinSpriteName = "#SIMULTANEOUS_CONNECTION_GREEN_SEAMLESS"
	SkinSpriteNameSimultaneousConnectionBlueSeamless    SkinSpriteName = "#SIMULTANEOUS_CONNECTION_BLUE_SEAMLESS"
	SkinSpriteNameSimultaneousConnectionYellowSeamless  SkinSpriteName = "#SIMULTANEOUS_CONNECTION_YELLOW_SEAMLESS"
	SkinSpriteNameSimultaneousConnectionPurpleSeamless  SkinSpriteName = "#SIMULTANEOUS_CONNECTION_PURPLE_SEAMLESS"
	SkinSpriteNameSimultaneousConnectionCyanSeamless    SkinSpriteName = "#SIMULTANEOUS_CONNECTION_CYAN_SEAMLESS"
	SkinSpriteNameDirectionalMarkerNeutral              SkinSpriteName = "#DIRECTIONAL_MARKER_NEUTRAL"
	SkinSpriteNameDirectionalMarkerRed                  SkinSpriteName = "#DIRECTIONAL_MARKER_RED"
	SkinSpriteNameDirectionalMarkerGreen                SkinSpriteName = "#DIRECTIONAL_MARKER_GREEN"
	SkinSpriteNameDirectionalMarkerBlue                 SkinSpriteName = "#DIRECTIONAL_MARKER_BLUE"
	SkinSpriteNameDirectionalMarkerYellow               SkinSpriteName = "#DIRECTIONAL_MARKER_YELLOW"
	SkinSpriteNameDirectionalMarkerPurple               SkinSpriteName = "#DIRECTIONAL_MARKER_PURPLE"
	SkinSpriteNameDirectionalMarkerCyan                 SkinSpriteName = "#DIRECTIONAL_MARKER_CYAN"
	SkinSpriteNameSimultaneousMarkerNeutral             SkinSpriteName = "#SIMULTANEOUS_MARKER_NEUTRAL"
	SkinSpriteNameSimultaneousMarkerRed                 SkinSpriteName = "#SIMULTANEOUS_MARKER_RED"
	SkinSpriteNameSimultaneousMarkerGreen               SkinSpriteName = "#SIMULTANEOUS_MARKER_GREEN"
	SkinSpriteNameSimultaneousMarkerBlue                SkinSpriteName = "#SIMULTANEOUS_MARKER_BLUE"
	SkinSpriteNameSimultaneousMarkerYellow              SkinSpriteName = "#SIMULTANEOUS_MARKER_YELLOW"
	SkinSpriteNameSimultaneousMarkerPurple              SkinSpriteName = "#SIMULTANEOUS_MARKER_PURPLE"
	SkinSpriteNameSimultaneousMarkerCyan                SkinSpriteName = "#SIMULTANEOUS_MARKER_CYAN"
	SkinSpriteNameStageMiddle                           SkinSpriteName = "#STAGE_MIDDLE"
	SkinSpriteNameStageLeftBorder                       SkinSpriteName = "#STAGE_LEFT_BORDER"
	SkinSpriteNameStageRightBorder                      SkinSpriteName = "#STAGE_RIGHT_BORDER"
	SkinSpriteNameStageTopBorder                        SkinSpriteName = "#STAGE_TOP_BORDER"
	SkinSpriteNameStageBottomBorder                     SkinSpriteName = "#STAGE_BOTTOM_BORDER"
	SkinSpriteNameStageLeftBorderSeamless               SkinSpriteName = "#STAGE_LEFT_BORDER_SEAMLESS"
	SkinSpriteNameStageRightBorderSeamless              SkinSpriteName = "#STAGE_RIGHT_BORDER_SEAMLESS"
	SkinSpriteNameStageTopBorderSeamless                SkinSpriteName = "#STAGE_TOP_BORDER_SEAMLESS"
	SkinSpriteNameStageBottomBorderSeamless             SkinSpriteName = "#STAGE_BOTTOM_BORDER_SEAMLESS"
	SkinSpriteNameStageTopLeftCorner                    SkinSpriteName = "#STAGE_TOP_LEFT_CORNER"
	SkinSpriteNameStageTopRightCorner                   SkinSpriteName = "#STAGE_TOP_RIGHT_CORNER"
	SkinSpriteNameStageBottomLeftCorner                 SkinSpriteName = "#STAGE_BOTTOM_LEFT_CORNER"
	SkinSpriteNameStageBottomRightCorner                SkinSpriteName = "#STAGE_BOTTOM_RIGHT_CORNER"
	SkinSpriteNameLane                                  SkinSpriteName = "#LANE"
	SkinSpriteNameLaneSeamless                          SkinSpriteName = "#LANE_SEAMLESS"
	SkinSpriteNameLaneAlternative                       SkinSpriteName = "#LANE_ALTERNATIVE"
	SkinSpriteNameLaneAlternativeSeamless               SkinSpriteName = "#LANE_ALTERNATIVE_SEAMLESS"
	SkinSpriteNameJudgmentLine                          SkinSpriteName = "#JUDGMENT_LINE"
	SkinSpriteNameNoteSlot                              SkinSpriteName = "#NOTE_SLOT"
	SkinSpriteNameStageCover                            SkinSpriteName = "#STAGE_COVER"
	SkinSpriteNameGridNeutral                           SkinSpriteName = "#GRID_NEUTRAL"
	SkinSpriteNameGridRed                               SkinSpriteName = "#GRID_RED"
	SkinSpriteNameGridGreen                             SkinSpriteName = "#GRID_GREEN"
	SkinSpriteNameGridBlue                              SkinSpriteName = "#GRID_BLUE"
	SkinSpriteNameGridYellow                            SkinSpriteName = "#GRID_YELLOW"
	SkinSpriteNameGridPurple                            SkinSpriteName = "#GRID_PURPLE"
	SkinSpriteNameGridCyan                              SkinSpriteName = "#GRID_CYAN"
)

type SkinItem struct {
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

type SkinData struct {
	Width         float64          `json:"width"`
	Height        float64          `json:"height"`
	Interpolation bool             `json:"interpolation"`
	Sprites       []SkinDataSprite `json:"sprites"`
}

type SkinDataSprite struct {
	Name      SkinSpriteName    `json:"name"`
	X         float64           `json:"x"`
	Y         float64           `json:"y"`
	W         float64           `json:"w"`
	H         float64           `json:"h"`
	Transform SkinDataTransform `json:"transform"`
}

type SkinDataTransform map[string]SkinDataExpression

type SkinDataExpression map[string]float64
