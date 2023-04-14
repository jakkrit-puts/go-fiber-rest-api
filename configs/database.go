package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/jakkrit-puts/go-fiber-rest-api/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connection() {
	dsn := os.Getenv("DB_DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Database: Connection fail., oops!.")
		log.Fatal(err)
	}

	fmt.Println("Database: Connected., let'go.")

	// Migrate Model
	db.AutoMigrate(&models.User{})

	DB = db
}
