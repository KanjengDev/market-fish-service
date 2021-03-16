package inventory

import "gorm.io/gorm"

type Repository interface {
	Save(inventory Inventory) (Inventory, error)
	FindAll() ([]Inventory, error)
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

func (r *repository) FindAll() ([]Inventory, error) {
	var inventory []Inventory

	err := r.db.Find(&inventory).Error
	if err != nil {
		return inventory, err
	}

	return inventory, nil
}
