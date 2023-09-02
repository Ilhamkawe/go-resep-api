package detailresep

type InputDetailResep struct {
	BahanID uint `json:"bahanID" binding:"required"`
	Jumlah  int  `json:"jumlah" binding:"required"`
}
