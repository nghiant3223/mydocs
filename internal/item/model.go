package item

import "github.com/nghiant3223/mydocs/pkg/basemodel"

type Item struct {
	basemodel.Model
	Type     Type    `json:"type"`
	Title    string  `json:"title"`
	Content  *string `json:"content,omitempty"`
	Priority int     `json:"priority"`
	Parent   *Item   `json:"-"`
	ParentID *uint   `json:"-"`
	Children []Item  `json:"children,omitempty"`
}
