package main

import (
	"database/sql"
	"fmt"
	"log"
	"market-fish-service/migration"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Database *gorm.DB
	Client   *sql.DB
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	//dsn local
	dsn := "root:@tcp(127.0.0.1:3306)/marketfish?charset=utf8mb4&parseTime=True&loc=Local"

	// dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ")/" + os.Getenv("DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	Database = db
	Client, _ = db.DB()
	autoCreate := os.Getenv("DB_MIGRATE")
	if autoCreate == "true" {
		fmt.Println("Drop and recreateing all tables....")
		migration.AutoMigrate(db)
		fmt.Println("All table recreated successfully ...")
	}

}
