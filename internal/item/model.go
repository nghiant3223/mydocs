package item

import "github.com/nghiant3223/mydocs/internal/basemodel"

type Item struct {
	basemodel.Model
	Title    string  `json:"title" gorm:"title"`
	Content  *string `json:"content" gorm:"content"`
	Order    int     `json:"order" gorm:"order"`
	Parent   *Item   `json:"-" gorm:"parent"`
	ParentID *uint   `json:"-" gorm:"parentID"`
	Children []Item  `json:"children"`
}
