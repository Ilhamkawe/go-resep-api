package bahan

type InputBahan struct {
	NamaBahan string `json:"namaBahan" binding:"required"`
	Satuan    string `json:"satuan" binding:"required"`
}

type InputBahanID struct {
	ID uint `uri:"id" binding:"required"`
}
