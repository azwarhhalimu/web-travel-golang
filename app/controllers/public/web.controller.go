package public

import "github.com/gofiber/fiber/v2"

func Index(ctx *fiber.Ctx) error {
	return ctx.Render("index",
		fiber.Map{
			"Nama":  "Azwar Halimu",
			"Title": "Beranda",
		},
		"layouts/public.template")
}
