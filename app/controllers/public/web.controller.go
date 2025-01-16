package public

import (
	"web_traveler/app/hook/render"
	"web_traveler/app/services/blog_services"
	"web_traveler/app/services/login"
	"web_traveler/app/services/lokasi_services"
	"web_traveler/app/services/user_services"
	"web_traveler/app/utils"

	"github.com/gofiber/fiber/v2"
)

func Index(ctx *fiber.Ctx) error {
	data := user_services.GetKategori()
	blog := blog_services.GetAll("USER")
	lokasi := lokasi_services.GetAll("USER")
	return render.RenderPublic(ctx, "index", fiber.Map{
		"Menu":         "index",
		"Data":         data,
		"Blog":         blog,
		"Lokasi":       lokasi,
		"NumberFormat": utils.NumberFormat,
	})
}
func SemuaLokasi(ctx *fiber.Ctx) error {
	lokasi := lokasi_services.GetAll()
	return render.RenderPublic(ctx, "semua-lokasi", fiber.Map{
		"NumberFormat": utils.NumberFormat,
		"Lokasi":       lokasi,
		"Title":        "Semua Lokasi",
		"Menu":         "semua-lokasi",
	})
}
func Blog(ctx *fiber.Ctx) error {
	blog := blog_services.GetAll()
	return render.RenderPublic(ctx, "blog", fiber.Map{
		"Blog":  blog,
		"Title": "Semua Blog",
		"Menu":  "blog",
	})
}
func LihatBlog(ctx *fiber.Ctx) error {
	alias := ctx.Params("alias")
	data, blog_lainnya := blog_services.Lihat_blog(alias)
	return render.RenderPublic(ctx, "lihat-blog", fiber.Map{
		"Data":         data,
		"Blog_lainnya": blog_lainnya,
	})
}

func LihatLokasi(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	data := lokasi_services.First(id, "USER")
	lokasi_lainnya := lokasi_services.LokasiLainnya(id)
	return render.RenderPublic(ctx, "lihat-lokasi", fiber.Map{
		"Data":           data,
		"Lokasi_lainnya": lokasi_lainnya,
		"NumberFormat":   utils.NumberFormat,
	})
}
func LokasiByKategori(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	data, kategori, list_kategori := lokasi_services.LokasiByKategori(id)
	return render.RenderPublic(ctx, "lokasi-by-kategori", fiber.Map{
		"Data":          data,
		"Kategori":      kategori,
		"List_kategori": list_kategori,
		"NumberFormat":  utils.NumberFormat,
	})
}

func Cari(ctx *fiber.Ctx) error {
	query := ctx.Query("query")
	data := lokasi_services.Cari(query)
	return render.RenderPublic(ctx, "cari", fiber.Map{
		"Data":         data,
		"NumberFormat": utils.NumberFormat,
		"Query":        query,
	})
}

func Login(ctx *fiber.Ctx) error {
	return ctx.Render("login", fiber.Map{})
}
func DoLogin(ctx *fiber.Ctx) error {
	username := ctx.FormValue("username")
	password := ctx.FormValue("password")
	status, _ := login.DoLogin(username, password)
	if status == "success" {
		return ctx.SendString("login sukses")
	} else {
		return ctx.SendString("login sukses status:" + status)
	}
}
