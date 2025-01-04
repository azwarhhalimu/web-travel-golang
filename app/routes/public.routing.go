package routes

import (
	"web_traveler/app/controllers/public"

	"github.com/gofiber/fiber/v2"
)

func PubliRouting(app *fiber.Group) {
	app.Get("/", public.Index).
		Get("/tentang", public.Tentang)
}
