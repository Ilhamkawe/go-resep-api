package kategori

import "go-resep-api/entity"

type KategoriFormatter struct {
	ID       uint   `json:"id"`
	Kategori string `json:"kategori"`
}

func FormatKategori(kategori entity.Kategori) KategoriFormatter {
	kategoriFormatter := KategoriFormatter{}
	kategoriFormatter.ID = kategori.ID
	kategoriFormatter.Kategori = kategori.Kategori

	return kategoriFormatter
}

func FormatsKategori(kategori []entity.Kategori) []KategoriFormatter {
	kategoriFormatters := []KategoriFormatter{}
	for _, kategori := range kategori {
		kategoriFormatter := FormatKategori(kategori)
		kategoriFormatters = append(kategoriFormatters, kategoriFormatter)
	}

	return kategoriFormatters
}
