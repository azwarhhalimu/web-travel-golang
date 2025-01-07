package kategori_services

import (
	"os"
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
func Delete(id_kategori string) {

	var data model.TblKategori
	config.DB.First(&data, id_kategori)
	//hapus
	config.DB.Delete(&model.TblKategori{}, id_kategori)
	path := "./storage/ikon/" + data.Ikon + ".svg"

	_, err := os.Stat(path)
	if err == nil {
		os.Remove(path)
	}
}
func GetFirst(id_kategori string) model.TblKategori {
	var data model.TblKategori
	config.DB.First(&data, id_kategori)
	return data
}

func Update(id_kategori string, kategori string) model.TblKategori {
	data := GetFirst(id_kategori)
	config.DB.
		Model(&model.TblKategori{}).
		Where("id_kategori", id_kategori).
		Updates(map[string]interface{}{
			"kategori": kategori,
		})
	return data
}
