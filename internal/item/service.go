package item

import "github.com/nghiant3223/mydocs/pkg/apperrors"

type Service interface {
	GetOneItem(id uint) (Item, error)
	GetItemTree() ([]Item, error)
	CreateItem(body CreateItemRequestBody) (Item, error)
	UpdateItem(id uint, body UpdateItemRequestBody) (Item, error)
	DeleteItem(id uint) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetOneItem(id uint) (Item, error) {
	return s.repo.FindByID(id)
}

func (s *service) GetItemTree() ([]Item, error) {
	parents, err := s.repo.FindByParentID(nil)
	if err != nil {
		return nil, err
	}
	var st stack
	for i := len(parents) - 1; i >= 0; i-- {
		st.push(&parents[i])
	}
	for !st.isEmpty() {
		top := st.pop()
		grandChildren, err := s.repo.FindByParentID(&(top.ID))
		if err != nil {
			return nil, err
		}

		top.Children = grandChildren
		for _, grandChild := range top.Children {
			grandChild := grandChild
			st.push(&grandChild)
		}
	}
	return parents, nil
}

func (s *service) CreateItem(body CreateItemRequestBody) (Item, error) {
	err := validateCreateItemRequestBody(body)
	if err != nil {
		return Item{}, err
	}
	item := body.toItem()
	return s.repo.Create(item)
}

func (s *service) UpdateItem(id uint, body UpdateItemRequestBody) (Item, error) {
	err := validateUpdateItemRequestBody(body)
	if err != nil {
		return Item{}, err
	}
	item := body.toItem()
	return s.repo.UpdateByID(id, item)
}

func (s *service) DeleteItem(id uint) error {
	return s.repo.DeleteByID(id)
}

func validateCreateItemRequestBody(body CreateItemRequestBody) error {
	if body.Type == Subject && body.Content != nil {
		return apperrors.InvalidItemData
	}
	if body.Type == Article && (body.Content == nil || *(body.Content) == "") {
		return apperrors.InvalidItemData
	}
	return nil
}

func validateUpdateItemRequestBody(body UpdateItemRequestBody) error {
	if body.Type == Subject && body.Content != nil {
		return apperrors.InvalidItemData
	}
	if body.Type == Article && (body.Content == nil || *(body.Content) == "") {
		return apperrors.InvalidItemData
	}
	return nil
}
