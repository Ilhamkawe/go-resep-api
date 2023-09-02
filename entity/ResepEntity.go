package entity

import "gorm.io/gorm"

type Resep struct {
	gorm.Model
	NamaResep  string   `gorm:"type:varchar(255);"`
	KategoriId uint     `gorm:"type:integer;not null"`
	Kategori   Kategori `gorm:"foreignKey:KategoriId;references:ID"`
}
