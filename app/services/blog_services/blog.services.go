package blog_services

import (
	"web_traveler/app/config"
	"web_traveler/app/model"

	"github.com/google/uuid"
)

func Save(formdata model.TblBlog) string {
	ui := uuid.New()
	namaimage := ui.String()
	config.DB.Table("blog").Create(map[string]interface{}{
		"judul":             formdata.Judul,
		"deskripsi_singkat": formdata.Deskripsi_Singkat,
		"deskripsi_lengkap": formdata.Deskripsi_Lengkap,
		"alias":             formdata.Alias,
		"nama_image":        namaimage,
	})
	return namaimage
}
func GetAll() []model.TblBlog {
	var data []model.TblBlog
	config.DB.Find(&data)
	return data
}
func First(id string) model.TblBlog {
	var data model.TblBlog
	config.DB.First(&data, id)
	return data
}
func Delete(id_blog string) model.TblBlog {
	data := First(id_blog)
	config.DB.Where("id_blog=?", id_blog).Delete(&model.TblBlog{})
	return data
}
func Update(id_blog string, formdata model.TblBlog) model.TblBlog {
	data := First(id_blog)
	config.DB.Model(&model.TblBlog{}).Where("id_blog=?", id_blog).Updates(map[string]interface{}{
		"judul":             formdata.Judul,
		"deskripsi_singkat": formdata.Deskripsi_Singkat,
		"deskripsi_lengkap": formdata.Deskripsi_Lengkap,
		"alias":             formdata.Alias,
	})
	return data
}
