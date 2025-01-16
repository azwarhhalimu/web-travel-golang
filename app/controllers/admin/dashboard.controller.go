package admin

import (
	"web_traveler/app/config"
	"web_traveler/app/hook/render"

	"github.com/gofiber/fiber/v2"
)

func Dashboard(ctx *fiber.Ctx) error {
	sesi, _ := config.STORE.Get(ctx)
	defer sesi.Save()
	login_success := sesi.Get("login_success")
	login_username := sesi.Get("login_username")
	login_name := sesi.Get("login_name")
	sesi.Delete("login_success")
	return render.RenderAdmin(ctx, "dashboard", fiber.Map{
		"Login_success":  login_success,
		"Login_username": login_username,
		"Login_name":     login_name,
	})
}

func Logout(ctx *fiber.Ctx) error {
	sesi, _ := config.STORE.Get(ctx)
	defer sesi.Save()
	sesi.Destroy()
	return ctx.Redirect("/login")
}
