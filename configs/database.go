package configs

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connection() {
	dsn := os.Getenv("DB_DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
	  fmt.Println("Database: Connection fail., oops!.")
    panic(err)
	}

	fmt.Println("Database: Connected., let'go.")
	DB = db
}