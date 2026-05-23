package resource

import (
	"github.com/WindowsSov8forUs/sonolus-core-go/core"
)

const PlaylistItemVersion = 1

type PlaylistItem struct {
	Name       string      `json:"name"`
	Source     string      `json:"source,omitempty"`
	Version    int         `json:"version"`
	Title      string      `json:"title"`
	Subtitle   string      `json:"subtitle"`
	Author     string      `json:"author"`
	AuthorUser *UserItem   `json:"authorUser,omitempty"`
	Tags       []core.Tag  `json:"tags"`
	Levels     []LevelItem `json:"levels"`
	Thumbnail  *core.Srl   `json:"thumbnail,omitempty"`
}

func DecodePlaylistItem(data []byte) (PlaylistItem, error) {
	return decodeItemAs[PlaylistItem](data)
}
