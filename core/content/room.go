package content

import (
	"github.com/WindowsSov8forUs/sonolus-core-go/core"
)

type RoomItem struct {
	Name       string     `json:"name"`
	Title      string     `json:"title"`
	Subtitle   string     `json:"subtitle"`
	Master     string     `json:"master"`
	MasterUser *UserItem  `json:"masterUser,omitempty"`
	Tags       []core.Tag `json:"tags"`
	Cover      *core.Srl  `json:"cover,omitempty"`
	BGM        *core.Srl  `json:"bgm,omitempty"`
	Preview    *core.Srl  `json:"preview,omitempty"`
}
