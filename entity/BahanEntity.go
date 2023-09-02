package entity

import "gorm.io/gorm"

type Bahan struct {
	gorm.Model
	NamaBahan string `gorm:"type:varchar(255);"`
	Satuan    string `gorm:"type:varchar(255);"`
}
