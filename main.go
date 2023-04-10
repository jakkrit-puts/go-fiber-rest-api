package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/jakkrit-puts/go-fiber-rest-api/routes"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")

	app := fiber.New()

	v1 := app.Group("/api/v1")

	routes.IndexRoutes(v1)
	routes.UserRoutes(v1)

	fmt.Printf("Port Server run is: %s", PORT)
	log.Fatal(app.Listen("localhost:" + PORT))

}
