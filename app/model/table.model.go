package model

type TblKategori struct {
	IDKategori uint `gorm:"primaryKey"`
	Kategori   string
	Ikon       string
}

func (TblKategori) TableName() string {
	return "kategori"
}
