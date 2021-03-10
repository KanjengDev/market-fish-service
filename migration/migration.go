package migration

import (
	"gorm.io/gorm"
)

var (
	Database *gorm.DB
)

type TableInterface interface {
	Pk() string
	Ref() string
	AddForeignkeys()
	InsertDefaults()
}

func AutoMigrate(database *gorm.DB) {
	Database = database

	// AutoMigrate
	database.AutoMigrate(
		Users{},
	)

}
