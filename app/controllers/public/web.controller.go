package public

import (
	"web_traveler/app/hook/render"

	"github.com/gofiber/fiber/v2"
)

func Index(ctx *fiber.Ctx) error {
	return ctx.Render("index",
		fiber.Map{
			"Nama":  "Azwar Halimu",
			"Title": "Beranda",
		},
		"layouts/public.template")
}
func Tentang(ctx *fiber.Ctx) error {
	return render.RenderPublic(ctx, "tentang", fiber.Map{
		"Title": "Tentang",
	})
}
