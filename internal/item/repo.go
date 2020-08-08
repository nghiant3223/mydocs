package item

import (
	"github.com/jinzhu/gorm"
)

type Repository interface {
	FindByID(id uint) (Item, error)
	FindByParentID(parentID *uint) ([]Item, error)
	Create(body CreateItemRequestBody) (Item, error)
	UpdateByID(id uint, body UpdateItemRequestBody) (Item, error)
	DeleteByID(id uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) FindByParentID(parentID *uint) ([]Item, error) {
	var items []Item
	var query *gorm.DB
	if parentID == nil {
		query = r.db.Where("parent_id IS NULL")
	} else {
		query = r.db.Where("parent_id = ?", parentID)
	}
	err := query.Order("'order' ASC").Find(&items).Error
	return items, err
}

func (r *repository) FindByID(id uint) (Item, error) {
	var item Item
	err := r.db.First(&item, id).Error
	return item, err
}

func (r *repository) Create(body CreateItemRequestBody) (Item, error) {
	item := body.toItem()
	err := r.db.Create(&item).Error
	return item, err
}

func (r *repository) UpdateByID(id uint, body UpdateItemRequestBody) (Item, error) {
	item, err := r.FindByID(id)
	if err != nil {
		return Item{}, err
	}
	err = r.db.Model(&item).Update(body).Error
	return item, err
}

func (r *repository) DeleteByID(id uint) error {
	return r.db.Where("id = ?", id).Delete(&Item{}).Error
}
