package inventory

import "gorm.io/gorm"

type Repository interface {
	Save(inventory Inventory) (Inventory, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(inventory Inventory) (Inventory, error) {
	err := r.db.Create(&inventory).Error
	if err != nil {
		return inventory, err
	}

	return inventory, nil
}
