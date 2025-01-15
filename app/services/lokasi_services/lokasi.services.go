package lokasi_services

import (
	"os"
	"strconv"
	"web_traveler/app/config"
	"web_traveler/app/model"
	"web_traveler/app/services/kategori_services"
)

func GetAll(flag ...string) []model.TblLokasi {

	var data []model.TblLokasi
	query := config.DB.
		Preload("Kategori").
		Preload("FotoLokasi", "default_foto=?", 1)

	if len(flag) > 0 {
		query.Limit(6).Order("RAND()")
	}
	query.Find(&data)

	return data
}
func First(id string, user ...string) model.TblLokasi {
	var data model.TblLokasi
	if len(user) > 0 {
		config.DB.Preload("Kategori").
			Preload("FotoLokasi").
			First(&data, id)
	} else {
		config.DB.Preload("Kategori").First(&data, id)
	}

	return data
}
func LokasiLainnya(id string) []model.TblLokasi {
	var data []model.TblLokasi
	config.DB.
		Limit(6).
		Preload("Kategori").
		Preload("FotoLokasi", "default_foto=?", "1").
		Find(&data)
	return data
}
func LokasiByKategori(id_kategori string) (lokasi []model.TblLokasi,
	xkategori model.TblKategori,
	list_kategori []model.TblKategori) {
	var data []model.TblLokasi
	var kategori model.TblKategori

	config.DB.First(&kategori, id_kategori)
	config.DB.
		Preload("Kategori").
		Preload("FotoLokasi", "default_foto=?", "1").
		Where("id_kategori=?", id_kategori).Find(&data)
	list_kategori = kategori_services.GetAll()
	lokasi = data
	xkategori = kategori
	return
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
func SimpanFotoLokasi(id string, nama_foto string, formdata model.TblFotoLokasi) string {
	config.DB.Table("foto_lokasi").Create(map[string]interface{}{
		"label":      formdata.Label,
		"image_name": nama_foto,
		"id_lokasi":  id,
	})
	return "success"
}
func GetFotoLokasi(id string) []model.TblFotoLokasi {
	var data []model.TblFotoLokasi
	config.DB.Where("id_lokasi", id).Find(&data)
	return data
}

func DeleteFoto(id string) string {
	var data model.TblFotoLokasi
	config.DB.First(&data, id)
	config.DB.Delete(&model.TblFotoLokasi{}, id)
	path := "./storage/foto-lokasi/" + data.Image_name + ".jpg"
	_, err := os.Stat(path)
	if err == nil {
		os.Remove(path)
	}
	return strconv.FormatUint(uint64(data.IDLokasi), 10)
}

func SetDefault(id_foto_lokasi string) string {
	var data model.TblFotoLokasi
	config.DB.First(&data, id_foto_lokasi)

	config.DB.Model(&model.TblFotoLokasi{}).
		Where("id_lokasi=?", data.IDLokasi).Update("default_foto", 0)

	config.DB.Model(&model.TblFotoLokasi{}).
		Where("id_foto_lokasi=?", id_foto_lokasi).Update("default_foto", 1)
	return strconv.FormatUint(uint64(data.IDLokasi), 10)
}
