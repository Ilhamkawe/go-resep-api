package resep

import (
	"go-resep-api/entity"
	"go-resep-api/src/detailresep"
)

type ResepFormatter struct {
	ID          uint                               `json:"id"`
	NamaResep   string                             `json:"namaResep"`
	Kategori    string                             `json:"kategori"`
	DetailResep []detailresep.ResepDetailFormatter `json:"detailResep"`
}

func FormatResep(resep entity.Resep) ResepFormatter {
	resepFormatter := ResepFormatter{}
	resepFormatter.ID = resep.ID
	resepFormatter.NamaResep = resep.NamaResep
	resepFormatter.Kategori = resep.Kategori.Kategori
	resepFormatter.DetailResep = detailresep.FormatDetailReseps(resep.DetailResep)
	return resepFormatter
}

func FormatReseps(resep []entity.Resep) []ResepFormatter {
	resepFormatters := []ResepFormatter{}
	for _, data := range resep {
		resepFormatter := FormatResep(data)
		resepFormatters = append(resepFormatters, resepFormatter)
	}

	return resepFormatters
}
