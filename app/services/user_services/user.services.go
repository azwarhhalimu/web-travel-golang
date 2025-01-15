package user_services

import (
	"web_traveler/app/config"
	"web_traveler/app/model"
)

func GetKategori() []model.TblKategori {
	var data []model.TblKategori
	config.DB.Table("kategori").
		Select("kategori.kategori, kategori.id_kategori,kategori.ikon, count(lokasi.id_lokasi) as jumlah_lokasi").
		Joins("LEFT JOIN lokasi ON lokasi.id_kategori=kategori.id_kategori").
		Group("kategori.id_kategori").
		Scan(&data)
	return data
}
