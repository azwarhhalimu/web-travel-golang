package model

type TblKategori struct {
	IDKategori uint `gorm:"primaryKey"`
	Kategori   string
	Ikon       string
}
type TblBlog struct {
	IDBlog            uint `gorm:"primaryKey"`
	Judul             string
	Deskripsi_Singkat string
	Deskripsi_Lengkap string
	Nama_Image        string
	Alias             string
}
type TblLokasi struct {
	IDLokasi       uint `gorm:"primaryKey"`
	IDKategori     int  `gorm:"foreignKey" form:"id_kategori"`
	Nama_lokasi    string
	Harga          uint
	Deskripsi      string
	Lat            string
	Lng            string
	Alamat_lengkap string
	Kategori       TblKategori `gorm:"foreignKey:IDKategori;references:IDKategori"`
}

func (TblKategori) TableName() string {
	return "kategori"
}
func (TblLokasi) TableName() string {
	return "lokasi"
}
func (TblBlog) TableName() string {
	return "blog"
}
