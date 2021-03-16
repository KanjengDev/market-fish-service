package migration

import "gorm.io/gorm"

type Inventory struct {
	gorm.Model
	Name        string `gorm:"default:null"`
	Price       int64  `gorm:"default:null"`
	FileName    string `gorm:"default:null"`
	Description string `sql:"type:text;"`
	Stock       int    `gorm:"default:0"`
	UserID      uint   `gorm:"default:null"`
}

func (Inventory) TableName() string {
	return "inventories"
}

func (Inventory) Pk() string {
	return "id"
}

func (I Inventory) Ref() string {
	return I.TableName() + "(" + I.Pk() + ")"
}
