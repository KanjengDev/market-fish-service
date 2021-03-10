package migration

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Name         string `gorm:"default:null"`
	Username     string `gorm:"default:null"`
	Phone        string `gorm:"default:null"`
	PasswordHash string `gorm:"default:null"`
	Address      string `gorm:"default:null"`
	RoleId       uint   `gorm:"not null"`
	Token        string `gorm:"default:null"`
}

func (Users) TableName() string {
	return "users"
}

func (Users) Pk() string {
	return "id"
}

func (u Users) Ref() string {
	return u.TableName() + "(" + u.Pk() + ")"
}
