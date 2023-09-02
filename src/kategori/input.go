package kategori

type InputKategori struct {
	Kategori string `json:"kategori" binding:"required"`
}

type InputKategoriID struct {
	ID uint `uri:"id" binding:"required"`
}
