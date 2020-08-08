package item

type stack []*Item

func (s *stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *stack) push(item *Item) {
	*s = append(*s, item)
}

func (s *stack) top() *Item {
	return (*s)[len(*s)-1]
}

func (s *stack) pop() *Item {
	index := len(*s) - 1
	element := (*s)[index]
	*s = (*s)[:index]
	return element
}
