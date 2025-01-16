package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func AuthAdmin(ctx *fiber.Ctx) error {
	print("auth middleware")
	return ctx.Next()
}
