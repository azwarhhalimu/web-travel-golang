package routes

import (
	"web_traveler/app/controllers/public"

	"github.com/gofiber/fiber/v2"
)

func PubliRouting(app *fiber.Group) {
	app.Get("/", public.Index).
		Get("/semua-lokasi.html", public.SemuaLokasi).
		Get("/blog.html", public.Blog).
		Get("/artikel/:alias", public.LihatBlog).
		Get("/lokasi/:id", public.LihatLokasi).
		Get("/lokasi-by-kategori/:id", public.LokasiByKategori).
		Get("/cari", public.Cari).
		Get("/login", public.Login)
}
