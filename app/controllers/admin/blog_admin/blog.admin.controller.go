package blog_admin

import (
	"fmt"
	"os"
	"web_traveler/app/hook/render"
	"web_traveler/app/model"
	"web_traveler/app/services/blog_services"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetAll(ctx *fiber.Ctx) error {
	data := blog_services.GetAll()
	ui := uuid.New().String()
	return render.RenderAdmin(ctx, "blog", fiber.Map{
		"Data": data,
		"Uuid": ui,
	})
}
func TambahBlog(ctx *fiber.Ctx) error {
	return render.RenderAdmin(ctx, "blog/tambah-blog", fiber.Map{
		"Summernote": true,
	})
}
func EditBlog(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	data := blog_services.First(id)
	return render.RenderAdmin(ctx, "blog/edit-blog", fiber.Map{
		"Summernote": true,
		"Data":       data,
	})
}
func SimpanBLog(ctx *fiber.Ctx) error {

	var data model.TblBlog
	ctx.BodyParser(&data)
	simpan := blog_services.Save(data)
	file_upload, err := ctx.FormFile("thumb")
	if err == nil {
		ctx.SaveFile(file_upload, "./storage/blog/"+simpan+".jpg")
	}
	return ctx.Redirect("/admin/blog.html")

}
func DeleteBlog(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	data := blog_services.Delete(id)
	path := "./storage/blog/" + data.Nama_Image + ".jpg"
	fmt.Println(path)
	_, err := os.Stat(path)
	if err == nil {
		os.Remove(path)
	}
	return ctx.Redirect("/admin/blog.html")
}
func UpdateBlog(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var data model.TblBlog
	ctx.BodyParser(&data)
	update := blog_services.Update(id, data)

	formupload, err := ctx.FormFile("thumb")
	if err == nil {
		ctx.SaveFile(formupload, "./storage/blog/"+update.Nama_Image+".jpg")
	}
	return ctx.Redirect("/admin/blog.html")
}
