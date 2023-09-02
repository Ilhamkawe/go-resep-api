package detailresep

import "go-resep-api/entity"

type ResepDetailFormatter struct {
	Bahan  string `json:"bahan"`
	Jumlah int    `json:"jumlah"`
	Satuan string `json:"satuan"`
}

func FormatDetailResep(detail entity.DetailResep) ResepDetailFormatter {
	resepDetailFormatter := ResepDetailFormatter{}
	resepDetailFormatter.Bahan = detail.Bahan.NamaBahan
	resepDetailFormatter.Jumlah = detail.Jumlah
	resepDetailFormatter.Satuan = detail.Bahan.Satuan

	return resepDetailFormatter
}

func FormatDetailReseps(details []entity.DetailResep) []ResepDetailFormatter {
	resepDetailFormatters := []ResepDetailFormatter{}
	for _, detail := range details {
		resepDetailFormatter := FormatDetailResep(detail)
		resepDetailFormatters = append(resepDetailFormatters, resepDetailFormatter)
	}

	return resepDetailFormatters
}
