package routes

import (
	"web_traveler/app/controllers/admin"
	"web_traveler/app/controllers/admin/blog_admin"
	"web_traveler/app/controllers/admin/kategori_admin"
	"web_traveler/app/controllers/admin/lokasi_admin"

	"github.com/gofiber/fiber/v2"
)

func AdminRouting(app *fiber.Group) {
	app.Get("/", admin.Dashboard)

	app.Get("/kategori.html", kategori_admin.Kategori)
	app.Get("/kategori/tambah-kategori.html", kategori_admin.TambahKategori)
	app.Post("/kategori/tambah-kategori.html", kategori_admin.SimpanKategori)
	app.Get("/kategori/delete/:id", kategori_admin.DeleteKategori)
	app.Get("/kategori/edit/:id", kategori_admin.EditKategori)
	app.Post("/kategori/edit/:id", kategori_admin.UpdateKategori)

	app.Get("/blog.html", blog_admin.GetAll)
	app.Get("/blog/tambah-blog.html", blog_admin.TambahBlog)
	app.Post("/blog/tambah-blog.html", blog_admin.SimpanBLog)
	app.Get("/blog/delete/:id", blog_admin.DeleteBlog) //delete data
	app.Get("/blog/edit/:id", blog_admin.EditBlog)     //edit data
	app.Post("/blog/edit/:id", blog_admin.UpdateBlog)

	app.Get("/lokasi.html", lokasi_admin.GetAll)                      //edit data
	app.Get("/lokasi/tambah-lokasi.html", lokasi_admin.TambahLokasi)  //tambah lokasi
	app.Post("/lokasi/tambah-lokasi.html", lokasi_admin.SimpanLokasi) //SIMPNA lokasi
	app.Get("/lokasi/delete/:id", lokasi_admin.DeleteLokasi)
	app.Get("/lokasi/edit-lokasi/:id", lokasi_admin.EditLokasi)        // DELETE LOKASI lokasi
	app.Post("/lokasi/edit-lokasi/:id", lokasi_admin.UpdateLokasi)     // UPDATE LOKASI
	app.Get("/lokasi/lihat-lokasi/:id", lokasi_admin.Lihat_lokasi)     // LIhat Lokasi LOKASI
	app.Get("/lokasi/foto-lokasi/:id", lokasi_admin.FotoLokasi)        // FOTO LOKASI
	app.Post("/lokasi/foto-lokasi/:id", lokasi_admin.SimpanFotoLokasi) //SIMPAN FOTO LOKASI
	app.Get("/foto-lokasi/delete/:id", lokasi_admin.HapusFotoLokasi)   //delete FOTO LOKASI
	app.Get("/foto-lokasi/default/:id", lokasi_admin.SetDefault)       //DefaultFOTO LOKASI

}
