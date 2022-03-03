package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
	"service-customer/database"
	"service-customer/routes"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	database.ConnectDB()

	routes.SetupRoutes(app)

	//log.Println("Start development server localhost:2022")
	log.Println("Start development server http://165.22.55.132:8005")

	//log.Fatal(app.Listen(":2022"))
	log.Fatal(app.Listen(":8005"))

}
