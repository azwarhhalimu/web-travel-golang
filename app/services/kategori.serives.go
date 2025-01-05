package kategori_services

import (
	"web_traveler/app/config"
	"web_traveler/app/model"
)

func GetAll() []model.TblKategori {
	var data []model.TblKategori
	config.DB.Find(&data)
	return data
}
func SimpanKategori(kategori string, nama_file string) bool {
	simpan := config.DB.Table("kategori").Create(map[string]interface{}{
		"kategori": kategori,
		"ikon":     nama_file,
	})
	if simpan.RowsAffected > 0 {
		return true
	}
	return false
}
func DeleteKategori(id string) any {
	simpan := config.DB.Delete(&model.TblKategori{}, "id_kategori=?", id)
	return simpan.RowsAffected
}
