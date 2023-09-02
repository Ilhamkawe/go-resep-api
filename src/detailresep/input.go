package detailresep

type InputDetailResep struct {
	ResepID uint `json:"resepID"`
	BahanID uint `json:"bahanID" binding:"required"`
	Jumlah  int  `json:"jumlah" binding:"required"`
}

type InputDetailID struct {
	ID uint `uri:"id" binding:"required"`
}
