package admin

import (
	"web_traveler/app/hook/render"

	"github.com/gofiber/fiber/v2"
)

func Dashboard(ctx *fiber.Ctx) error {
	return render.RenderAdmin(ctx, "dashboard")
}
