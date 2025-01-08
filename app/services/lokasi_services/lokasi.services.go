package lokasi_services

import (
	"web_traveler/app/config"
	"web_traveler/app/model"
)

func GetAll() []model.TblLokasi {
	var data []model.TblLokasi
	config.DB.Find(&data)
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
