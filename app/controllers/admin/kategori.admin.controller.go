package admin_kategori

import (
	"web_traveler/app/hook/render"
	kategori_services "web_traveler/app/services"

	"github.com/gofiber/fiber/v2"
)

func Kategori(ctx *fiber.Ctx) error {
	data := kategori_services.GetAll()
	return render.RenderAdmin(ctx, "kategori", fiber.Map{
		"Data": data,
	})
}
