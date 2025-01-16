package middleware

import (
	"web_traveler/app/config"

	"github.com/gofiber/fiber/v2"
)

func AuthAdmin(ctx *fiber.Ctx) error {
	sesi, _ := config.STORE.Get(ctx)

	login_username := sesi.Get("login_username")
	if login_username == nil {
		sesi.Set("login_error", "Anda harus login terlebih dahulu")
		defer sesi.Save()
		return ctx.Redirect("/login")
	}
	return ctx.Next()
}
