package item

type UpdateItemRequestBody struct {
	Title   string  `json:"title"`
	Content *string `json:"content"`
	Order   int     `json:"order"`
}

type CreateItemRequestBody struct {
	Title   string  `json:"title"`
	Content *string `json:"content"`
	Order   int     `json:"order"`
}

func (b *CreateItemRequestBody) toItem() Item {
	return Item{
		Title:   b.Title,
		Content: b.Content,
		Order:   b.Order,
	}
}
