package lokasi_services

import (
	"web_traveler/app/config"
	"web_traveler/app/model"
)

func GetAll() []model.TblLokasi {
	var data []model.TblLokasi
	config.DB.Preload("Kategori").Find(&data)
	return data
}
func Fist(id string) model.TblLokasi {
	var data model.TblLokasi
	config.DB.Preload("Kategori").First(&data, id)
	return data
}
func Delete(id string) string {
	config.DB.Delete(&model.TblLokasi{}, id)
	return "success"
}
func Save(formdata model.TblLokasi) string {
	config.DB.Table("lokasi").Create(map[string]interface{}{

		"id_kategori":    formdata.IDKategori,
		"nama_lokasi":    formdata.Nama_lokasi,
		"harga":          formdata.Harga,
		"deskripsi":      formdata.Deskripsi,
		"lat":            formdata.Lat,
		"lng":            formdata.Lng,
		"alamat_lengkap": formdata.Alamat_lengkap,
	})
	return "success"
}
func Update(id string, formdata model.TblLokasi) string {
	config.DB.Model(&model.TblLokasi{}).
		Where("id_lokasi=?", id).
		Updates(map[string]interface{}{
			"nama_lokasi":    formdata.Nama_lokasi,
			"harga":          formdata.Harga,
			"deskripsi":      formdata.Deskripsi,
			"lat":            formdata.Lat,
			"lng":            formdata.Lng,
			"alamat_lengkap": formdata.Alamat_lengkap,
		})
	return "success"
}
