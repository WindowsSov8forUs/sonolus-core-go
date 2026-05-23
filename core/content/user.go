package content

import "github.com/WindowsSov8forUs/sonolus-core-go/core"

type UserItem struct {
	Name   string     `json:"name"`
	Source string     `json:"source,omitempty"`
	Title  string     `json:"title"`
	Handle string     `json:"handle,omitempty"`
	Tags   []core.Tag `json:"tags,omitempty"`
}
