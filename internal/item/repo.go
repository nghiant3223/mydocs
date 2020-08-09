package item

import (
	"github.com/jinzhu/gorm"
	"github.com/nghiant3223/mydocs/pkg/apperrors"
)

type Repository interface {
	FindByID(id uint) (Item, error)
	FindByParentID(parentID *uint) ([]Item, error)
	Create(item Item) (Item, error)
	UpdateByID(id uint, item Item) (Item, error)
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
	err := query.Order("priority ASC").Find(&items).Error
	return items, err
}

func (r *repository) FindByID(id uint) (Item, error) {
	var item Item
	err := r.db.First(&item, id).Error
	if gorm.IsRecordNotFoundError(err) {
		err = apperrors.RecordNotFound
	}
	return item, err
}

func (r *repository) Create(item Item) (Item, error) {
	err := r.db.Create(&item).Error
	return item, err
}

func (r *repository) UpdateByID(id uint, updates Item) (Item, error) {
	item, err := r.FindByID(id)
	if err != nil {
		return Item{}, err
	}
	err = r.db.Model(&item).Update(updates).Error
	return item, err
}

func (r *repository) DeleteByID(id uint) error {
	return r.db.Where("id = ?", id).Unscoped().Delete(&Item{}).Error
}
