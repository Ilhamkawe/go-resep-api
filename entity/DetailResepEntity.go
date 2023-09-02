package entity

import "gorm.io/gorm"

type DetailResep struct {
	gorm.Model
	ResepId uint  `gorm:"type:integer;not null"`
	Resep   Resep `gorm:"foreignKey:ResepId;references:ID"`
	BahanId uint  `gorm:"type:integer;not null"`
	Bahan   Bahan `gorm:"foreignKey:BahanId;references:ID"`
	Jumlah  int   `gorm:"type:integer;column:'jumlah'"`
}
