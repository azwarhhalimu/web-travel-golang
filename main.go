package main

import (
	"log"
	"web_traveler/app/config"
	"web_traveler/app/middleware"
	"web_traveler/app/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/jet/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		println(".ENV GAGAL DI LOAD. FILE .ENV TIDAK DI TEMUKAN")
	}
	config.InitSession()
	config.Connect()
	template := jet.New("./app/views", ".jet")

	app := fiber.New(fiber.Config{
		Views: template,
	})
	app.Static("storage", "./assets")
	app.Static("images", "./storage")

	public := app.Group("/").(*fiber.Group)
	routes.PubliRouting(public)

	admin := app.Group("/admin", middleware.AuthAdmin).(*fiber.Group)
	routes.AdminRouting(admin)

	log.Panic(app.Listen(":2000"))

}
