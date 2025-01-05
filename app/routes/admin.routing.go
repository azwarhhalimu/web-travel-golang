package routes

import (
	admin_kategori "web_traveler/app/controllers/admin"

	"github.com/gofiber/fiber/v2"
)

func AdminRouting(app *fiber.Group) {
	app.Get("/kategori.html", admin_kategori.Kategori)
	app.Get("/kategori/tambah-kategori.html", admin_kategori.TambahKategori)
}
