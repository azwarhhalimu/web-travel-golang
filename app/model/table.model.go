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
	Kategori       TblKategori     `gorm:"foreignKey:IDKategori;references:IDKategori"`
	FotoLokasi     []TblFotoLokasi `gorm:"foreignKey:IDLokasi"`
}
type TblFotoLokasi struct {
	IDFotoLokasi uint `gorm:"primaryKey"`
	IDLokasi     uint `gorm:"foreignKey"`
	Label        string
	Default_foto uint
	Image_name   string
}

func (TblFotoLokasi) TableName() string {
	return "foto_lokasi"
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
