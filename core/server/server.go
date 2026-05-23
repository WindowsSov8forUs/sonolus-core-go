package server

import (
	"github.com/WindowsSov8forUs/sonolus-core-go/core"
	"github.com/WindowsSov8forUs/sonolus-core-go/core/content"
)

type ServerAuthenticateRequest struct {
	Type        string             `json:"type"`
	Address     string             `json:"address"`
	Time        float64            `json:"time"`
	UserProfile ServiceUserProfile `json:"userProfile"`
}

type ServerAuthenticateResponse struct {
	Session    string  `json:"session"`
	Expiration float64 `json:"expiration"`
}

type ServerMessage struct {
	Message string `json:"message,omitempty"`
}

type ServerConfiguration struct {
	Options []ServerOption `json:"options"`
}

type ServerForm struct {
	Type                string         `json:"type"`
	Title               core.Text      `json:"title"`
	Icon                core.Icon      `json:"icon,omitempty"`
	Description         string         `json:"description,omitempty"`
	Help                string         `json:"help,omitempty"`
	RequireConfirmation bool           `json:"requireConfirmation"`
	Options             []ServerOption `json:"options"`
}

type ServerOption interface {
	serverOption()
	OptionType() string
}

type ServerOptionBase struct {
	Query       string    `json:"query"`
	Name        core.Text `json:"name"`
	Description string    `json:"description,omitempty"`
	Required    bool      `json:"required"`
}

func (ServerOptionBase) serverOption() {}

type ServerTextOption struct {
	ServerOptionBase
	Type        string   `json:"type"`
	Def         string   `json:"def"`
	Placeholder string   `json:"placeholder"`
	Limit       int      `json:"limit"`
	Shortcuts   []string `json:"shortcuts"`
}

func (ServerTextOption) OptionType() string { return "text" }

type ServerTextAreaOption struct {
	ServerOptionBase
	Type        string   `json:"type"`
	Def         string   `json:"def"`
	Placeholder string   `json:"placeholder"`
	Limit       int      `json:"limit"`
	Shortcuts   []string `json:"shortcuts"`
}

func (ServerTextAreaOption) OptionType() string { return "textArea" }

type ServerSliderOption struct {
	ServerOptionBase
	Type string    `json:"type"`
	Def  float64   `json:"def"`
	Min  float64   `json:"min"`
	Max  float64   `json:"max"`
	Step float64   `json:"step"`
	Unit core.Text `json:"unit,omitempty"`
}

func (ServerSliderOption) OptionType() string { return "slider" }

type ServerToggleOption struct {
	ServerOptionBase
	Type string `json:"type"`
	Def  bool   `json:"def"`
}

func (ServerToggleOption) OptionType() string { return "toggle" }

type ServerSelectOption struct {
	ServerOptionBase
	Type   string                    `json:"type"`
	Def    string                    `json:"def"`
	Values []ServerSelectOptionValue `json:"values"`
}

func (ServerSelectOption) OptionType() string { return "select" }

type ServerSelectOptionValue struct {
	Name  string    `json:"name"`
	Title core.Text `json:"title"`
}

type ServerMultiOption struct {
	ServerOptionBase
	Type   string                    `json:"type"`
	Def    []bool                    `json:"def"`
	Values []ServerSelectOptionValue `json:"values"`
}

func (ServerMultiOption) OptionType() string { return "multi" }

type ServerServerItemOption struct {
	ServerOptionBase
	Type              string        `json:"type"`
	ItemType          core.ItemType `json:"itemType"`
	InfoType          string        `json:"infoType,omitempty"`
	Def               *core.Sil     `json:"def"`
	AllowOtherServers bool          `json:"allowOtherServers"`
}

func (ServerServerItemOption) OptionType() string { return "serverItem" }

type ServerServerItemsOption struct {
	ServerOptionBase
	Type              string        `json:"type"`
	ItemType          core.ItemType `json:"itemType"`
	InfoType          string        `json:"infoType,omitempty"`
	Def               []core.Sil    `json:"def"`
	AllowOtherServers bool          `json:"allowOtherServers"`
	Limit             int           `json:"limit"`
}

func (ServerServerItemsOption) OptionType() string { return "serverItems" }

type ServerCollectionItemOption struct {
	ServerOptionBase
	Type     string        `json:"type"`
	ItemType core.ItemType `json:"itemType"`
}

func (ServerCollectionItemOption) OptionType() string { return "collectionItem" }

type ServerFileOption struct {
	ServerOptionBase
	Type       string                      `json:"type"`
	Def        string                      `json:"def"`
	Validation *ServerFileOptionValidation `json:"validation,omitempty"`
}

func (ServerFileOption) OptionType() string { return "file" }

type ServerFileOptionValidation struct {
	Type      string   `json:"type"`
	MinSize   *float64 `json:"minSize,omitempty"`
	MaxSize   *float64 `json:"maxSize,omitempty"`
	MinWidth  *float64 `json:"minWidth,omitempty"`
	MaxWidth  *float64 `json:"maxWidth,omitempty"`
	MinHeight *float64 `json:"minHeight,omitempty"`
	MaxHeight *float64 `json:"maxHeight,omitempty"`
	MinLength *float64 `json:"minLength,omitempty"`
	MaxLength *float64 `json:"maxLength,omitempty"`
}

type ServerFileOptionValidationFile = ServerFileOptionValidation

type ServerFileOptionValidationImage = ServerFileOptionValidation

type ServerFileOptionValidationAudio = ServerFileOptionValidation

type ServerFileOptionValidationZip = ServerFileOptionValidation

type ServerFileOptionValidationJson = ServerFileOptionValidation

type ServerInfo struct {
	Title         string              `json:"title"`
	Description   string              `json:"description,omitempty"`
	Buttons       []ServerInfoButton  `json:"buttons"`
	Configuration ServerConfiguration `json:"configuration"`
	Banner        *core.Srl           `json:"banner,omitempty"`
}

type ServerInfoButton interface {
	serverInfoButton()
	ButtonType() string
}

type ServerInfoAuthenticationButton struct {
	Type string `json:"type"`
}

func (ServerInfoAuthenticationButton) serverInfoButton()  {}
func (ServerInfoAuthenticationButton) ButtonType() string { return "authentication" }

type ServerInfoItemButton struct {
	Type       core.ItemType `json:"type"`
	Title      core.Text     `json:"title,omitempty"`
	Icon       core.Icon     `json:"icon,omitempty"`
	BadgeCount *float64      `json:"badgeCount,omitempty"`
	InfoType   string        `json:"infoType,omitempty"`
	ItemName   string        `json:"itemName,omitempty"`
}

func (ServerInfoItemButton) serverInfoButton()  {}
func (ServerInfoItemButton) ButtonType() string { return string(ServerInfoItemButton{}.Type) }

type ServerInfoConfigurationButton struct {
	Type string `json:"type"`
}

func (ServerInfoConfigurationButton) serverInfoButton()  {}
func (ServerInfoConfigurationButton) ButtonType() string { return "configuration" }

type ServerCreateItemRequest struct {
	Values string `json:"values"`
}

type ServerCreateItemResponse struct {
	Key                  string   `json:"key"`
	Hashes               []string `json:"hashes"`
	ShouldUpdateInfo     *bool    `json:"shouldUpdateInfo,omitempty"`
	ShouldNavigateToItem string   `json:"shouldNavigateToItem,omitempty"`
}

type ServerItemDetails[T any] struct {
	Item         T                       `json:"item"`
	Description  string                  `json:"description,omitempty"`
	Actions      []ServerForm            `json:"actions"`
	HasCommunity bool                    `json:"hasCommunity"`
	Leaderboards []ServerItemLeaderboard `json:"leaderboards"`
	Sections     []ServerItemSection     `json:"sections"`
}

type ServerItemInfo struct {
	Title             core.Text           `json:"title,omitempty"`
	Creates           []ServerForm        `json:"creates,omitempty"`
	Searches          []ServerForm        `json:"searches,omitempty"`
	QuickSearchValues string              `json:"quickSearchValues,omitempty"`
	Sections          []ServerItemSection `json:"sections"`
	Banner            *core.Srl           `json:"banner,omitempty"`
}

type ServerItemList[T any] struct {
	Title             core.Text    `json:"title,omitempty"`
	PageCount         int          `json:"pageCount"`
	Cursor            string       `json:"cursor,omitempty"`
	Items             []T          `json:"items"`
	Searches          []ServerForm `json:"searches,omitempty"`
	QuickSearchValues string       `json:"quickSearchValues,omitempty"`
}

type ServerItemSection struct {
	Title        core.Text     `json:"title"`
	Icon         core.Icon     `json:"icon,omitempty"`
	Description  string        `json:"description,omitempty"`
	Help         string        `json:"help,omitempty"`
	ItemType     core.ItemType `json:"itemType"`
	Items        []any         `json:"items"`
	Search       *ServerForm   `json:"search,omitempty"`
	SearchValues string        `json:"searchValues,omitempty"`
}

type ServerItemSectionTyped[T any] struct {
	Title        core.Text     `json:"title"`
	Icon         core.Icon     `json:"icon,omitempty"`
	Description  string        `json:"description,omitempty"`
	Help         string        `json:"help,omitempty"`
	ItemType     core.ItemType `json:"itemType"`
	Items        []T           `json:"items"`
	Search       *ServerForm   `json:"search,omitempty"`
	SearchValues string        `json:"searchValues,omitempty"`
}

type ServerSubmitItemActionRequest struct {
	Values string `json:"values"`
}

type ServerSubmitItemActionResponse struct {
	Key                  string   `json:"key"`
	Hashes               []string `json:"hashes"`
	ShouldUpdateItem     *bool    `json:"shouldUpdateItem,omitempty"`
	ShouldRemoveItem     *bool    `json:"shouldRemoveItem,omitempty"`
	ShouldNavigateToItem string   `json:"shouldNavigateToItem,omitempty"`
}

type ServerUploadItemResponse struct {
	ShouldUpdateInfo     *bool  `json:"shouldUpdateInfo,omitempty"`
	ShouldNavigateToItem string `json:"shouldNavigateToItem,omitempty"`
}

type ServerUploadItemActionResponse struct {
	ShouldUpdateItem     *bool  `json:"shouldUpdateItem,omitempty"`
	ShouldRemoveItem     *bool  `json:"shouldRemoveItem,omitempty"`
	ShouldNavigateToItem string `json:"shouldNavigateToItem,omitempty"`
}

type ServerLevelResultInfo struct {
	Submits []ServerForm `json:"submits,omitempty"`
}

type ServerSubmitLevelResultRequest struct {
	Replay content.ReplayItem `json:"replay"`
	Values string             `json:"values"`
}

type ServerSubmitLevelResultResponse struct {
	Key    string   `json:"key"`
	Hashes []string `json:"hashes"`
}

type ServerUploadLevelResultResponse struct{}

type ServerItemLeaderboard struct {
	Name        string    `json:"name"`
	Title       core.Text `json:"title"`
	Description string    `json:"description,omitempty"`
}

type ServerItemLeaderboardDetails struct {
	TopRecords []ServerItemLeaderboardRecord `json:"topRecords"`
}

type ServerItemLeaderboardRecord struct {
	Name       string            `json:"name"`
	Rank       core.Text         `json:"rank"`
	Player     string            `json:"player"`
	PlayerUser *content.UserItem `json:"playerUser,omitempty"`
	Value      core.Text         `json:"value"`
}

type ServerItemLeaderboardRecordList struct {
	PageCount int                           `json:"pageCount"`
	Cursor    string                        `json:"cursor,omitempty"`
	Records   []ServerItemLeaderboardRecord `json:"records"`
}

type ServerItemLeaderboardRecordDetails struct {
	Replays []content.ReplayItem `json:"replays"`
}

type ServerItemCommunityComment struct {
	Name       string            `json:"name"`
	Author     string            `json:"author"`
	AuthorUser *content.UserItem `json:"authorUser,omitempty"`
	Time       float64           `json:"time"`
	Content    string            `json:"content"`
	Actions    []ServerForm      `json:"actions"`
}

type ServerItemCommunityCommentList struct {
	PageCount int                          `json:"pageCount"`
	Cursor    string                       `json:"cursor,omitempty"`
	Comments  []ServerItemCommunityComment `json:"comments"`
}

type ServerSubmitItemCommunityCommentActionRequest struct {
	Values string `json:"values"`
}

type ServerSubmitItemCommunityCommentActionResponse struct {
	Key                          string   `json:"key"`
	Hashes                       []string `json:"hashes"`
	ShouldUpdateCommunity        *bool    `json:"shouldUpdateCommunity,omitempty"`
	ShouldUpdateComments         *bool    `json:"shouldUpdateComments,omitempty"`
	ShouldNavigateCommentsToPage *float64 `json:"shouldNavigateCommentsToPage,omitempty"`
}

type ServerUploadItemCommunityCommentActionResponse struct {
	ShouldUpdateCommunity        *bool    `json:"shouldUpdateCommunity,omitempty"`
	ShouldUpdateComments         *bool    `json:"shouldUpdateComments,omitempty"`
	ShouldNavigateCommentsToPage *float64 `json:"shouldNavigateCommentsToPage,omitempty"`
}

type ServerItemCommunityInfo struct {
	Actions     []ServerForm                 `json:"actions"`
	TopComments []ServerItemCommunityComment `json:"topComments"`
}

type ServerSubmitItemCommunityActionRequest struct {
	Values string `json:"values"`
}

type ServerSubmitItemCommunityActionResponse struct {
	Key                          string   `json:"key"`
	Hashes                       []string `json:"hashes"`
	ShouldUpdateCommunity        *bool    `json:"shouldUpdateCommunity,omitempty"`
	ShouldUpdateComments         *bool    `json:"shouldUpdateComments,omitempty"`
	ShouldNavigateCommentsToPage *float64 `json:"shouldNavigateCommentsToPage,omitempty"`
}

type ServerUploadItemCommunityActionResponse struct {
	ShouldUpdateCommunity        *bool    `json:"shouldUpdateCommunity,omitempty"`
	ShouldUpdateComments         *bool    `json:"shouldUpdateComments,omitempty"`
	ShouldNavigateCommentsToPage *float64 `json:"shouldNavigateCommentsToPage,omitempty"`
}

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
	Type  string            `json:"type"`
	State string            `json:"state"`
	Level content.LevelItem `json:"level"`
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
