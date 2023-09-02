package entity

import "gorm.io/gorm"

type Kategori struct {
	gorm.Model
	Kategori string `gorm:"type:varchar(255);"`
}
