package server

import (
	"encoding/json"

	"github.com/WindowsSov8forUs/sonolus-core-go/core"
	"github.com/WindowsSov8forUs/sonolus-core-go/core/resource"
	coreerrors "github.com/WindowsSov8forUs/sonolus-core-go/errors"
)

type ChatMessage interface{ chatMessage() }

type QuickChatMessage struct {
	UserID ServiceUserId `json:"userId"`
	Type   string        `json:"type"`
	Value  string        `json:"value"`
}

func (QuickChatMessage) chatMessage() {}

type TextChatMessage struct {
	UserID *ServiceUserId `json:"userId"`
	Type   string         `json:"type"`
	Value  string         `json:"value"`
}

func (TextChatMessage) chatMessage() {}

type ServerCreateRoomRequest struct{}

type ServerCreateRoomResponse struct {
	Name string `json:"name"`
	Key  string `json:"key"`
}

type ServerJoinRoomRequest struct {
	Type        string             `json:"type"`
	Address     string             `json:"address"`
	Room        string             `json:"room"`
	Time        float64            `json:"time"`
	UserProfile ServiceUserProfile `json:"userProfile"`
}

type ServerJoinRoomResponse struct {
	URL     string `json:"url"`
	Type    string `json:"type"`
	Session string `json:"session"`
}

type LevelOptionEntry struct {
	Index int     `json:"index"`
	Value float64 `json:"value"`
}

type ResultEntry struct {
	UserID ServiceUserId       `json:"userId"`
	Result core.GameplayResult `json:"result"`
}

type RoomStatus string

const (
	RoomStatusSelecting RoomStatus = "selecting"
	RoomStatusPreparing RoomStatus = "preparing"
	RoomStatusPlaying   RoomStatus = "playing"
)

type RoomUser struct {
	Authentication string `json:"authentication"`
	Signature      string `json:"signature"`
}

type ScoreboardSection struct {
	Title  core.Text    `json:"title"`
	Icon   core.Icon    `json:"icon,omitempty"`
	Scores []ScoreEntry `json:"scores"`
}

type ScoreEntry struct {
	UserID ServiceUserId `json:"userId"`
	Value  core.Text     `json:"value"`
}

type Suggestion struct {
	UserID ServiceUserId `json:"userId"`
	Level  core.Sil      `json:"level"`
}

type UserStatus string

const (
	UserStatusWaiting UserStatus = "waiting"
	UserStatusReady   UserStatus = "ready"
	UserStatusSkipped UserStatus = "skipped"
	UserStatusPlaying UserStatus = "playing"
)

type UserStatusEntry struct {
	UserID ServiceUserId `json:"userId"`
	Status UserStatus    `json:"status"`
}

type Command interface{ command() }

type AddChatMessageCommand struct {
	Type    string      `json:"type"`
	Nonce   float64     `json:"nonce"`
	Message ChatMessage `json:"message"`
}

func (AddChatMessageCommand) command() {}

type AddSuggestionCommand struct {
	Type  string   `json:"type"`
	Level core.Sil `json:"level"`
}

func (AddSuggestionCommand) command() {}

type ClearSuggestionsCommand struct {
	Type string `json:"type"`
}

func (ClearSuggestionsCommand) command() {}

type FinishGameplayCommand struct {
	Type   string              `json:"type"`
	State  string              `json:"state"`
	Result core.GameplayResult `json:"result"`
}

func (FinishGameplayCommand) command() {}

type RemoveSuggestionCommand struct {
	Type       string     `json:"type"`
	Suggestion Suggestion `json:"suggestion"`
}

func (RemoveSuggestionCommand) command() {}

type RemoveUserCommand struct {
	Type   string        `json:"type"`
	UserID ServiceUserId `json:"userId"`
}

func (RemoveUserCommand) command() {}

type ReportUserCommand struct {
	Type         string        `json:"type"`
	UserID       ServiceUserId `json:"userId"`
	OptionValues string        `json:"optionValues"`
}

func (ReportUserCommand) command() {}

type ResetScoreboardCommand struct {
	Type string `json:"type"`
}

func (ResetScoreboardCommand) command() {}

type StartGameplayCommand struct {
	Type  string             `json:"type"`
	State string             `json:"state"`
	Level resource.LevelItem `json:"level"`
}

func (StartGameplayCommand) command() {}

type SwapSuggestionsCommand struct {
	Type        string     `json:"type"`
	SuggestionA Suggestion `json:"suggestionA"`
	SuggestionB Suggestion `json:"suggestionB"`
}

func (SwapSuggestionsCommand) command() {}

type UpdateAutoExitCommand struct {
	Type     string        `json:"type"`
	AutoExit core.AutoExit `json:"autoExit"`
}

func (UpdateAutoExitCommand) command() {}

type UpdateIsSuggestionsLockedCommand struct {
	Type                string `json:"type"`
	IsSuggestionsLocked bool   `json:"isSuggestionsLocked"`
}

func (UpdateIsSuggestionsLockedCommand) command() {}

type UpdateLeadCommand struct {
	Type string        `json:"type"`
	Lead ServiceUserId `json:"lead"`
}

func (UpdateLeadCommand) command() {}

type UpdateLevelCommand struct {
	Type  string   `json:"type"`
	Level core.Sil `json:"level"`
}

func (UpdateLevelCommand) command() {}

type UpdateLevelOptionCommand struct {
	Type        string           `json:"type"`
	LevelOption LevelOptionEntry `json:"levelOption"`
}

func (UpdateLevelOptionCommand) command() {}

type UpdateMasterCommand struct {
	Type   string        `json:"type"`
	Master ServiceUserId `json:"master"`
}

func (UpdateMasterCommand) command() {}

type UpdateOptionValuesCommand struct {
	Type         string `json:"type"`
	OptionValues string `json:"optionValues"`
}

func (UpdateOptionValuesCommand) command() {}

type UpdateStatusCommand struct {
	Type   string     `json:"type"`
	Status RoomStatus `json:"status"`
}

func (UpdateStatusCommand) command() {}

type UpdateUserStatusCommand struct {
	Type   string     `json:"type"`
	Status UserStatus `json:"status"`
}

func (UpdateUserStatusCommand) command() {}

type Event interface{ event() }

type AddChatMessageEvent struct {
	Type    string      `json:"type"`
	Nonce   *float64    `json:"nonce,omitempty"`
	Message ChatMessage `json:"message"`
}

func (AddChatMessageEvent) event() {}

type AddResultEvent struct {
	Type   string      `json:"type"`
	Result ResultEntry `json:"result"`
}

func (AddResultEvent) event() {}

type AddSuggestionEvent struct {
	Type       string     `json:"type"`
	Suggestion Suggestion `json:"suggestion"`
}

func (AddSuggestionEvent) event() {}

type AddUserEvent struct {
	Type string   `json:"type"`
	User RoomUser `json:"user"`
}

func (AddUserEvent) event() {}

type ArrangeScoreboardSectionScoresEvent struct {
	Type         string `json:"type"`
	SectionIndex int    `json:"sectionIndex"`
	Indexes      []int  `json:"indexes"`
}

func (ArrangeScoreboardSectionScoresEvent) event() {}

type ClearSuggestionsEvent struct {
	Type string `json:"type"`
}

func (ClearSuggestionsEvent) event() {}

type InsertScoreboardSectionEvent struct {
	Type              string            `json:"type"`
	Index             int               `json:"index"`
	ScoreboardSection ScoreboardSection `json:"scoreboardSection"`
}

func (InsertScoreboardSectionEvent) event() {}

type InsertScoreboardSectionScoreEvent struct {
	Type         string     `json:"type"`
	SectionIndex int        `json:"sectionIndex"`
	Index        int        `json:"index"`
	Score        ScoreEntry `json:"score"`
}

func (InsertScoreboardSectionScoreEvent) event() {}

type MergeScoreboardSectionScoresEvent struct {
	Type         string `json:"type"`
	SectionIndex int    `json:"sectionIndex"`
	Indexes      []int  `json:"indexes"`
}

func (MergeScoreboardSectionScoresEvent) event() {}

type MoveScoreboardSectionEvent struct {
	Type      string `json:"type"`
	FromIndex int    `json:"fromIndex"`
	ToIndex   int    `json:"toIndex"`
}

func (MoveScoreboardSectionEvent) event() {}

type MoveScoreboardSectionScoreEvent struct {
	Type         string `json:"type"`
	SectionIndex int    `json:"sectionIndex"`
	FromIndex    int    `json:"fromIndex"`
	ToIndex      int    `json:"toIndex"`
}

func (MoveScoreboardSectionScoreEvent) event() {}

type RemoveScoreboardSectionEvent struct {
	Type  string `json:"type"`
	Index int    `json:"index"`
}

func (RemoveScoreboardSectionEvent) event() {}

type RemoveScoreboardSectionScoreEvent struct {
	Type         string `json:"type"`
	SectionIndex int    `json:"sectionIndex"`
	Index        int    `json:"index"`
}

func (RemoveScoreboardSectionScoreEvent) event() {}

type RemoveSuggestionEvent struct {
	Type       string     `json:"type"`
	Suggestion Suggestion `json:"suggestion"`
}

func (RemoveSuggestionEvent) event() {}

type RemoveUserEvent struct {
	Type string   `json:"type"`
	User RoomUser `json:"user"`
}

func (RemoveUserEvent) event() {}

type StartRoundEvent struct {
	Type  string  `json:"type"`
	State string  `json:"state"`
	Seed  float64 `json:"seed"`
}

func (StartRoundEvent) event() {}

type SwapScoreboardSectionScoresEvent struct {
	Type         string `json:"type"`
	SectionIndex int    `json:"sectionIndex"`
	IndexA       int    `json:"indexA"`
	IndexB       int    `json:"indexB"`
}

func (SwapScoreboardSectionScoresEvent) event() {}

type SwapScoreboardSectionsEvent struct {
	Type   string `json:"type"`
	IndexA int    `json:"indexA"`
	IndexB int    `json:"indexB"`
}

func (SwapScoreboardSectionsEvent) event() {}

type SwapSuggestionsEvent struct {
	Type        string     `json:"type"`
	SuggestionA Suggestion `json:"suggestionA"`
	SuggestionB Suggestion `json:"suggestionB"`
}

func (SwapSuggestionsEvent) event() {}

type UpdateAutoExitEvent struct {
	Type     string        `json:"type"`
	AutoExit core.AutoExit `json:"autoExit"`
}

func (UpdateAutoExitEvent) event() {}

type UpdateEvent struct {
	Type                  string              `json:"type"`
	AllowOtherServers     bool                `json:"allowOtherServers"`
	ReportUserOptions     []ServerForm        `json:"reportUserOptions"`
	Title                 string              `json:"title"`
	Status                RoomStatus          `json:"status"`
	Master                *ServiceUserId      `json:"master"`
	Lead                  *ServiceUserId      `json:"lead"`
	Options               []ServerForm        `json:"options"`
	OptionValues          string              `json:"optionValues"`
	Level                 *core.Sil           `json:"level"`
	LevelOptions          []LevelOptionEntry  `json:"levelOptions"`
	AutoExit              core.AutoExit       `json:"autoExit"`
	IsSuggestionsLocked   bool                `json:"isSuggestionsLocked"`
	Suggestions           []Suggestion        `json:"suggestions"`
	ScoreboardDescription string              `json:"scoreboardDescription,omitempty"`
	ScoreboardSections    []ScoreboardSection `json:"scoreboardSections"`
	Results               []ResultEntry       `json:"results"`
	Users                 []RoomUser          `json:"users"`
	UserStatuses          []UserStatusEntry   `json:"userStatuses"`
}

func (UpdateEvent) event() {}

type UpdateIsSuggestionsLockedEvent struct {
	Type                string `json:"type"`
	IsSuggestionsLocked bool   `json:"isSuggestionsLocked"`
}

func (UpdateIsSuggestionsLockedEvent) event() {}

type UpdateLeadEvent struct {
	Type string         `json:"type"`
	Lead *ServiceUserId `json:"lead"`
}

func (UpdateLeadEvent) event() {}

type UpdateLevelEvent struct {
	Type  string    `json:"type"`
	Level *core.Sil `json:"level"`
}

func (UpdateLevelEvent) event() {}

type UpdateLevelOptionEvent struct {
	Type        string           `json:"type"`
	LevelOption LevelOptionEntry `json:"levelOption"`
}

func (UpdateLevelOptionEvent) event() {}

type UpdateLevelOptionsEvent struct {
	Type         string             `json:"type"`
	LevelOptions []LevelOptionEntry `json:"levelOptions"`
}

func (UpdateLevelOptionsEvent) event() {}

type UpdateMasterEvent struct {
	Type   string         `json:"type"`
	Master *ServiceUserId `json:"master"`
}

func (UpdateMasterEvent) event() {}

type UpdateOptionsEvent struct {
	Type         string       `json:"type"`
	Options      []ServerForm `json:"options"`
	OptionValues string       `json:"optionValues"`
}

func (UpdateOptionsEvent) event() {}

type UpdateOptionValuesEvent struct {
	Type         string `json:"type"`
	OptionValues string `json:"optionValues"`
}

func (UpdateOptionValuesEvent) event() {}

type UpdateScoreboardDescriptionEvent struct {
	Type                  string `json:"type"`
	ScoreboardDescription string `json:"scoreboardDescription,omitempty"`
}

func (UpdateScoreboardDescriptionEvent) event() {}

type UpdateScoreboardSectionEvent struct {
	Type              string            `json:"type"`
	Index             int               `json:"index"`
	ScoreboardSection ScoreboardSection `json:"scoreboardSection"`
}

func (UpdateScoreboardSectionEvent) event() {}

type UpdateScoreboardSectionIconEvent struct {
	Type  string    `json:"type"`
	Index int       `json:"index"`
	Icon  core.Icon `json:"icon,omitempty"`
}

func (UpdateScoreboardSectionIconEvent) event() {}

type UpdateScoreboardSectionScoresEvent struct {
	Type   string       `json:"type"`
	Index  int          `json:"index"`
	Scores []ScoreEntry `json:"scores"`
}

func (UpdateScoreboardSectionScoresEvent) event() {}

type UpdateScoreboardSectionScoresValueEvent struct {
	Type   string      `json:"type"`
	Index  int         `json:"index"`
	Values []core.Text `json:"values"`
}

func (UpdateScoreboardSectionScoresValueEvent) event() {}

type UpdateScoreboardSectionsEvent struct {
	Type               string              `json:"type"`
	ScoreboardSections []ScoreboardSection `json:"scoreboardSections"`
}

func (UpdateScoreboardSectionsEvent) event() {}

type UpdateScoreboardSectionTitleEvent struct {
	Type  string    `json:"type"`
	Index int       `json:"index"`
	Title core.Text `json:"title"`
}

func (UpdateScoreboardSectionTitleEvent) event() {}

type UpdateStatusEvent struct {
	Type   string     `json:"type"`
	Status RoomStatus `json:"status"`
}

func (UpdateStatusEvent) event() {}

type UpdateSuggestionsEvent struct {
	Type        string       `json:"type"`
	Suggestions []Suggestion `json:"suggestions"`
}

func (UpdateSuggestionsEvent) event() {}

type UpdateTitleEvent struct {
	Type  string `json:"type"`
	Title string `json:"title"`
}

func (UpdateTitleEvent) event() {}

type UpdateUsersEvent struct {
	Type  string     `json:"type"`
	Users []RoomUser `json:"users"`
}

func (UpdateUsersEvent) event() {}

type UpdateUserStatusesEvent struct {
	Type         string            `json:"type"`
	UserStatuses []UserStatusEntry `json:"userStatuses"`
}

func (UpdateUserStatusesEvent) event() {}

type UpdateUserStatusEvent struct {
	Type       string          `json:"type"`
	UserStatus UserStatusEntry `json:"userStatus"`
}

func (UpdateUserStatusEvent) event() {}

func DecodeChatMessage(data []byte) (ChatMessage, error) {
	var header unionTypeHeader
	if err := json.Unmarshal(data, &header); err != nil {
		return nil, err
	}

	switch header.Type {
	case "quick":
		var message QuickChatMessage
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	case "text":
		var message TextChatMessage
		if err := json.Unmarshal(data, &message); err != nil {
			return nil, err
		}
		return message, nil
	default:
		return nil, coreerrors.UnknownUnionTypeError{
			Union: "ChatMessage",
			Type:  header.Type,
		}
	}
}

func DecodeCommand(data []byte) (Command, error) {
	var header unionTypeHeader
	if err := json.Unmarshal(data, &header); err != nil {
		return nil, err
	}

	switch header.Type {
	case "addChatMessage":
		return decodeCommandAs[AddChatMessageCommand](data)
	case "addSuggestion":
		return decodeCommandAs[AddSuggestionCommand](data)
	case "clearSuggestions":
		return decodeCommandAs[ClearSuggestionsCommand](data)
	case "finishGameplay":
		return decodeCommandAs[FinishGameplayCommand](data)
	case "removeSuggestion":
		return decodeCommandAs[RemoveSuggestionCommand](data)
	case "removeUser":
		return decodeCommandAs[RemoveUserCommand](data)
	case "reportUser":
		return decodeCommandAs[ReportUserCommand](data)
	case "resetScoreboard":
		return decodeCommandAs[ResetScoreboardCommand](data)
	case "startGameplay":
		return decodeCommandAs[StartGameplayCommand](data)
	case "swapSuggestions":
		return decodeCommandAs[SwapSuggestionsCommand](data)
	case "updateAutoExit":
		return decodeCommandAs[UpdateAutoExitCommand](data)
	case "updateIsSuggestionsLocked":
		return decodeCommandAs[UpdateIsSuggestionsLockedCommand](data)
	case "updateLead":
		return decodeCommandAs[UpdateLeadCommand](data)
	case "updateLevel":
		return decodeCommandAs[UpdateLevelCommand](data)
	case "updateLevelOption":
		return decodeCommandAs[UpdateLevelOptionCommand](data)
	case "updateMaster":
		return decodeCommandAs[UpdateMasterCommand](data)
	case "updateOptionValues":
		return decodeCommandAs[UpdateOptionValuesCommand](data)
	case "updateStatus":
		return decodeCommandAs[UpdateStatusCommand](data)
	case "updateUserStatus":
		return decodeCommandAs[UpdateUserStatusCommand](data)
	default:
		return nil, coreerrors.UnknownUnionTypeError{
			Union: "Command",
			Type:  header.Type,
		}
	}
}

func decodeCommandAs[T Command](data []byte) (Command, error) {
	var command T
	if err := json.Unmarshal(data, &command); err != nil {
		return nil, err
	}
	return command, nil
}

func DecodeEvent(data []byte) (Event, error) {
	var header unionTypeHeader
	if err := json.Unmarshal(data, &header); err != nil {
		return nil, err
	}

	switch header.Type {
	case "addChatMessage":
		return decodeEventAs[AddChatMessageEvent](data)
	case "addResult":
		return decodeEventAs[AddResultEvent](data)
	case "addSuggestion":
		return decodeEventAs[AddSuggestionEvent](data)
	case "addUser":
		return decodeEventAs[AddUserEvent](data)
	case "arrangeScoreboardSectionScores":
		return decodeEventAs[ArrangeScoreboardSectionScoresEvent](data)
	case "clearSuggestions":
		return decodeEventAs[ClearSuggestionsEvent](data)
	case "insertScoreboardSection":
		return decodeEventAs[InsertScoreboardSectionEvent](data)
	case "insertScoreboardSectionScore":
		return decodeEventAs[InsertScoreboardSectionScoreEvent](data)
	case "mergeScoreboardSectionScores":
		return decodeEventAs[MergeScoreboardSectionScoresEvent](data)
	case "moveScoreboardSection":
		return decodeEventAs[MoveScoreboardSectionEvent](data)
	case "moveScoreboardSectionScore":
		return decodeEventAs[MoveScoreboardSectionScoreEvent](data)
	case "removeScoreboardSection":
		return decodeEventAs[RemoveScoreboardSectionEvent](data)
	case "removeScoreboardSectionScore":
		return decodeEventAs[RemoveScoreboardSectionScoreEvent](data)
	case "removeSuggestion":
		return decodeEventAs[RemoveSuggestionEvent](data)
	case "removeUser":
		return decodeEventAs[RemoveUserEvent](data)
	case "startRound":
		return decodeEventAs[StartRoundEvent](data)
	case "swapScoreboardSectionScores":
		return decodeEventAs[SwapScoreboardSectionScoresEvent](data)
	case "swapScoreboardSections":
		return decodeEventAs[SwapScoreboardSectionsEvent](data)
	case "swapSuggestions":
		return decodeEventAs[SwapSuggestionsEvent](data)
	case "updateAutoExit":
		return decodeEventAs[UpdateAutoExitEvent](data)
	case "update":
		return decodeEventAs[UpdateEvent](data)
	case "updateIsSuggestionsLocked":
		return decodeEventAs[UpdateIsSuggestionsLockedEvent](data)
	case "updateLead":
		return decodeEventAs[UpdateLeadEvent](data)
	case "updateLevel":
		return decodeEventAs[UpdateLevelEvent](data)
	case "updateLevelOption":
		return decodeEventAs[UpdateLevelOptionEvent](data)
	case "updateLevelOptions":
		return decodeEventAs[UpdateLevelOptionsEvent](data)
	case "updateMaster":
		return decodeEventAs[UpdateMasterEvent](data)
	case "updateOptions":
		return decodeEventAs[UpdateOptionsEvent](data)
	case "updateOptionValues":
		return decodeEventAs[UpdateOptionValuesEvent](data)
	case "updateScoreboardDescription":
		return decodeEventAs[UpdateScoreboardDescriptionEvent](data)
	case "updateScoreboardSection":
		return decodeEventAs[UpdateScoreboardSectionEvent](data)
	case "updateScoreboardSectionIcon":
		return decodeEventAs[UpdateScoreboardSectionIconEvent](data)
	case "updateScoreboardSectionScores":
		return decodeEventAs[UpdateScoreboardSectionScoresEvent](data)
	case "updateScoreboardSectionScoresValue":
		return decodeEventAs[UpdateScoreboardSectionScoresValueEvent](data)
	case "updateScoreboardSections":
		return decodeEventAs[UpdateScoreboardSectionsEvent](data)
	case "updateScoreboardSectionTitle":
		return decodeEventAs[UpdateScoreboardSectionTitleEvent](data)
	case "updateStatus":
		return decodeEventAs[UpdateStatusEvent](data)
	case "updateSuggestions":
		return decodeEventAs[UpdateSuggestionsEvent](data)
	case "updateTitle":
		return decodeEventAs[UpdateTitleEvent](data)
	case "updateUsers":
		return decodeEventAs[UpdateUsersEvent](data)
	case "updateUserStatuses":
		return decodeEventAs[UpdateUserStatusesEvent](data)
	case "updateUserStatus":
		return decodeEventAs[UpdateUserStatusEvent](data)
	default:
		return nil, coreerrors.UnknownUnionTypeError{
			Union: "Event",
			Type:  header.Type,
		}
	}
}

func decodeEventAs[T Event](data []byte) (Event, error) {
	var event T
	if err := json.Unmarshal(data, &event); err != nil {
		return nil, err
	}
	return event, nil
}

func (command *AddChatMessageCommand) UnmarshalJSON(data []byte) error {
	type addChatMessageCommand AddChatMessageCommand
	var raw struct {
		Message json.RawMessage `json:"message"`
		*addChatMessageCommand
	}
	raw.addChatMessageCommand = (*addChatMessageCommand)(command)
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	message, err := DecodeChatMessage(raw.Message)
	if err != nil {
		return err
	}
	command.Message = message
	return nil
}

func (event *AddChatMessageEvent) UnmarshalJSON(data []byte) error {
	type addChatMessageEvent AddChatMessageEvent
	var raw struct {
		Message json.RawMessage `json:"message"`
		*addChatMessageEvent
	}
	raw.addChatMessageEvent = (*addChatMessageEvent)(event)
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	message, err := DecodeChatMessage(raw.Message)
	if err != nil {
		return err
	}
	event.Message = message
	return nil
}
