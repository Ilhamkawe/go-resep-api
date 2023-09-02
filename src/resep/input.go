package resep

import (
	"go-resep-api/src/detailresep"
)

type InputResep struct {
	NamaResep   string                         `json:"namaResep" binding:"required"`
	KategoriID  uint                           `json:"kategoriID" binding:"required"`
	DetailResep []detailresep.InputDetailResep `json:"detailResep" binding:"required"`
}

type InputResepID struct {
	ID uint `uri:"id" binding:"required"`
}
