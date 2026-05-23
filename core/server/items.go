package server

import (
	"encoding/json"

	"github.com/WindowsSov8forUs/sonolus-core-go/core"
	"github.com/WindowsSov8forUs/sonolus-core-go/core/resource"
)

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

func (section *ServerItemSection) UnmarshalJSON(data []byte) error {
	type serverItemSection ServerItemSection
	var raw struct {
		Items []json.RawMessage `json:"items"`
		*serverItemSection
	}
	raw.serverItemSection = (*serverItemSection)(section)
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	items := make([]any, len(raw.Items))
	for i, rawItem := range raw.Items {
		item, err := resource.DecodeItem(section.ItemType, rawItem)
		if err != nil {
			return err
		}
		items[i] = item
	}
	section.Items = items
	return nil
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
	Replay resource.ReplayItem `json:"replay"`
	Values string              `json:"values"`
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
	Name       string             `json:"name"`
	Rank       core.Text          `json:"rank"`
	Player     string             `json:"player"`
	PlayerUser *resource.UserItem `json:"playerUser,omitempty"`
	Value      core.Text          `json:"value"`
}

type ServerItemLeaderboardRecordList struct {
	PageCount int                           `json:"pageCount"`
	Cursor    string                        `json:"cursor,omitempty"`
	Records   []ServerItemLeaderboardRecord `json:"records"`
}

type ServerItemLeaderboardRecordDetails struct {
	Replays []resource.ReplayItem `json:"replays"`
}

type ServerItemCommunityComment struct {
	Name       string             `json:"name"`
	Author     string             `json:"author"`
	AuthorUser *resource.UserItem `json:"authorUser,omitempty"`
	Time       float64            `json:"time"`
	Content    string             `json:"content"`
	Actions    []ServerForm       `json:"actions"`
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
