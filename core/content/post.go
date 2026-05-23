package content

import (
	"github.com/WindowsSov8forUs/sonolus-core-go/core"
)

const PostItemVersion = 1

type PostItem struct {
	Name       string     `json:"name"`
	Source     string     `json:"source,omitempty"`
	Version    int        `json:"version"`
	Title      string     `json:"title"`
	Time       float64    `json:"time"`
	Author     string     `json:"author"`
	AuthorUser *UserItem  `json:"authorUser,omitempty"`
	Tags       []core.Tag `json:"tags"`
	Thumbnail  *core.Srl  `json:"thumbnail,omitempty"`
}
