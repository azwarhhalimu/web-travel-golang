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
func GetAll(flag ...string) []model.TblBlog {
	var data []model.TblBlog
	if len(flag) > 0 {
		config.DB.Limit(3).Order("RAND()").Find(&data)
	} else {
		config.DB.Find(&data)
	}

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

func Lihat_blog(alias string) (lihat_blog model.TblBlog, blog_lainnya []model.TblBlog) {
	var data_blog_lainnya []model.TblBlog
	var data model.TblBlog
	config.DB.Where("alias=?", alias).First(&data)
	config.DB.Where("alias!=?", alias).Limit(3).Find(&data_blog_lainnya)

	lihat_blog = data
	blog_lainnya = data_blog_lainnya
	return
}
