package kategori_admin

import (
	"web_traveler/app/hook/render"
	"web_traveler/app/services/kategori_services"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func Kategori(ctx *fiber.Ctx) error {
	data := kategori_services.GetAll()
	random := uuid.New()
	return render.RenderAdmin(ctx, "kategori", fiber.Map{
		"Data":   data,
		"Random": random.String(),
	})
}
func TambahKategori(ctx *fiber.Ctx) error {
	return render.RenderAdmin(ctx, "kategori/tambah-kategori")
}
func SimpanKategori(ctx *fiber.Ctx) error {

	//UPLOAD
	name := uuid.New()
	file_image, _ := ctx.FormFile("ikon")
	upload := ctx.SaveFile(file_image, "./storage/ikon/"+name.String()+".svg")
	if upload == nil {
		kategori := ctx.FormValue("kategori")
		simpan := kategori_services.SimpanKategori(kategori, name.String())
		if simpan {
			ctx.Redirect("/admin/kategori.html")
		}
	}
	return ctx.SendString("data berhasil di simpan")

}

func DeleteKategori(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	kategori_services.Delete(id)
	return ctx.Redirect("/admin/kategori.html")
}
func EditKategori(ctx *fiber.Ctx) error {

	id := ctx.Params("id")
	data := kategori_services.GetFirst(id)
	return render.RenderAdmin(ctx, "kategori/edit-kategori",
		fiber.Map{
			"Data": data,
		},
	)
}
func UpdateKategori(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	kategori := ctx.FormValue("kategori")
	data := kategori_services.Update(id, kategori)
	path := "./storage/ikon/" + data.Ikon + ".svg"
	dataform, err := ctx.FormFile("ikon")
	if err == nil {
		ctx.SaveFile(dataform, path)
	}
	return ctx.Redirect("/admin/kategori.html")
}
