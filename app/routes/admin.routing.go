package routes

import (
	"web_traveler/app/controllers/admin"

	"github.com/gofiber/fiber/v2"
)

func AdminRouting(app *fiber.Group) {
	app.Get("/", admin.Dashboard).
		Get("/kategori.html", admin.Kategori).
		Get("/blog.html", admin.Blog).
		Get("/lokasi.html", admin.Lokasi)
}
