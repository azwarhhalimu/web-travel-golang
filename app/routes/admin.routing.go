package routes

import (
	"web_traveler/app/controllers/admin/kategori_admin"

	"github.com/gofiber/fiber/v2"
)

func AdminRouting(app *fiber.Group) {
	app.Get("/kategori.html", kategori_admin.Kategori)
	app.Get("/kategori/tambah-kategori.html", kategori_admin.TambahKategori)
	app.Post("/kategori/tambah-kategori.html", kategori_admin.SimpanKategori)
	app.Get("/kategori/delete/:id", kategori_admin.DeleteKategori)

	app.Get("/kategori/edit/:id", kategori_admin.EditKategori)
	app.Post("/kategori/edit/:id", kategori_admin.UpdateKategori)

}
