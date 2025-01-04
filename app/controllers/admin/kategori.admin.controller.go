package admin_kategori

import (
	"web_traveler/app/hook/render"

	"github.com/gofiber/fiber/v2"
)

func Kategori(ctx *fiber.Ctx) error {
	return render.RenderAdmin(ctx, "kategori")
}
