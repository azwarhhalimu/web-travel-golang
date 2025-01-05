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
