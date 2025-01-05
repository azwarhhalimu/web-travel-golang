package admin_kategori

import (
	"web_traveler/app/hook/render"
	kategori_services "web_traveler/app/services"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func Kategori(ctx *fiber.Ctx) error {
	data := kategori_services.GetAll()
	return render.RenderAdmin(ctx, "kategori", fiber.Map{
		"Data": data,
	})
}
func TambahKategori(ctx *fiber.Ctx) error {
	return render.RenderAdmin(ctx, "kategori/tambah-kategori")
}
func SimpanKategori(ctx *fiber.Ctx) error {

	//UPLOAD
	name := uuid.New()
	file_image, _ := ctx.FormFile("ikon")

	ctx.SaveFile(file_image, "./storage/ikon/"+name.String()+".svg")
	return ctx.SendString("data berhasil di simpan")
}
