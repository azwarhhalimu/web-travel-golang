package lokasi_admin

import (
	"web_traveler/app/hook/render"
	"web_traveler/app/model"
	"web_traveler/app/services/kategori_services"
	"web_traveler/app/services/lokasi_services"

	"github.com/gofiber/fiber/v2"
)

func GetAll(ctx *fiber.Ctx) error {
	data := lokasi_services.GetAll()
	return render.RenderAdmin(ctx, "lokasi", fiber.Map{
		"Data": data,
	})
}
func TambahLokasi(ctx *fiber.Ctx) error {
	data := kategori_services.GetAll()
	return render.RenderAdmin(ctx, "lokasi/tambah-lokasi", fiber.Map{
		"Summernote": true,
		"Kategori":   data,
	})
}
func SimpanLokasi(ctx *fiber.Ctx) error {
	var data model.TblLokasi
	ctx.BodyParser(&data)
	lokasi_services.Save(data)
	return ctx.Redirect("/admin/lokasi.html")
}

func DeleteLokasi(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	lokasi_services.Delete(id)
	return ctx.Redirect("/admin/lokasi.html")
}
