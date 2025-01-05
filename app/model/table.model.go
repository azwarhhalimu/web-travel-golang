package model

type TblKategori struct {
	IDKategori uint `gorm:"primaryKey"`
	Kategori   string
	Icon       string
}

func (TblKategori) TableName() string {
	return "kategori"
}
