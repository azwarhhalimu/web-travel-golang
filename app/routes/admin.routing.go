package routes

import (
	"web_traveler/app/controllers/admin"

	"github.com/gofiber/fiber/v2"
)

func AdminRouting(app *fiber.Group) {
	app.Get("/blog", admin.GetBlog)
}
