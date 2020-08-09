package item

type UpdateItemRequestBody struct {
	Type    Type    `json:"type"`
	Title   string  `json:"title"`
	Content *string `json:"content"`
	Order   int     `json:"order"`
}

func (b *UpdateItemRequestBody) toItem() Item {
	return Item{
		Title:    b.Title,
		Content:  b.Content,
		Priority: b.Order,
		Type:     b.Type,
	}
}

type CreateItemRequestBody struct {
	Type    Type    `json:"type"`
	Title   string  `json:"title"`
	Content *string `json:"content"`
	Order   int     `json:"order"`
}

func (b *CreateItemRequestBody) toItem() Item {
	return Item{
		Title:    b.Title,
		Content:  b.Content,
		Priority: b.Order,
		Type:     b.Type,
	}
}
