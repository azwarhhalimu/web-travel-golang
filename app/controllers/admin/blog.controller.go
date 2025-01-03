package admin

import "github.com/gofiber/fiber/v2"

func GetBlog(ctx *fiber.Ctx) error {
	return ctx.SendString("ini halaman blog")
}
