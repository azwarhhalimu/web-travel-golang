package lokasi_admin

import (
	"web_traveler/app/hook/render"
	"web_traveler/app/model"
	"web_traveler/app/services/kategori_services"
	"web_traveler/app/services/lokasi_services"
	"web_traveler/app/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetAll(ctx *fiber.Ctx) error {
	data := lokasi_services.GetAll()
	return render.RenderAdmin(ctx, "lokasi", fiber.Map{
		"Data":         data,
		"NumberFormat": utils.NumberFormat,
	})
}
func TambahLokasi(ctx *fiber.Ctx) error {

	kategori := kategori_services.GetAll()
	return render.RenderAdmin(ctx, "lokasi/tambah-lokasi", fiber.Map{
		"Summernote": true,
		"Kategori":   kategori,
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

func EditLokasi(ctx *fiber.Ctx) error {
	kategori := kategori_services.GetAll()
	id := ctx.Params("id")
	data := lokasi_services.Fist(id)
	return render.RenderAdmin(ctx, "lokasi/edit-lokasi", fiber.Map{
		"Summernote": true,
		"Kategori":   kategori,
		"Data":       data,
	})
}
func UpdateLokasi(ctx *fiber.Ctx) error {
	var data model.TblLokasi
	ctx.BodyParser(&data)
	id := ctx.Params("id")

	lokasi_services.Update(id, data)
	return ctx.Redirect("/admin/lokasi.html")
}
func Lihat_lokasi(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	data := lokasi_services.Fist(id)
	return render.RenderAdmin(ctx, "lokasi/lihat-lokasi", fiber.Map{
		"Data": data,
	})
}
func FotoLokasi(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	data := lokasi_services.GetFotoLokasi(id)
	return render.RenderAdmin(ctx, "lokasi/foto-lokasi", fiber.Map{
		"Data": data,
	})
}
func SimpanFotoLokasi(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	foto := uuid.New().String()
	var data model.TblFotoLokasi
	ctx.BodyParser(&data)
	simpan := lokasi_services.SimpanFotoLokasi(id, foto, data)
	if simpan == "success" {
		file, err := ctx.FormFile("foto")
		if err == nil {
			ctx.SaveFile(file, "./storage/foto-lokasi/"+foto+".jpg")
		}

	}
	return ctx.Redirect("/admin/lokasi/foto-lokasi/" + id)

}
func HapusFotoLokasi(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	delete := lokasi_services.DeleteFoto(id)
	return ctx.Redirect("/admin/lokasi/foto-lokasi/" + delete)
}

func SetDefault(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	ubah := lokasi_services.SetDefault(id)
	return ctx.Redirect("/admin/lokasi/foto-lokasi/" + ubah)
}
