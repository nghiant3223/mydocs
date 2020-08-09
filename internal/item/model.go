package item

import "github.com/nghiant3223/mydocs/pkg/basemodel"

type Item struct {
	basemodel.Model
	Type     Type    `json:"type" gorm:"type"`
	Title    string  `json:"title" gorm:"title"`
	Content  *string `json:"content,omitempty" gorm:"content"`
	Priority int     `json:"priority" gorm:"priority"`
	Parent   *Item   `json:"-" gorm:"parent"`
	ParentID *uint   `json:"-" gorm:"parentID"`
	Children []Item  `json:"children,omitempty"`
}
