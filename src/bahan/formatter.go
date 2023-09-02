package bahan

import "go-resep-api/entity"

type BahanFormatter struct {
	NamaBahan string `json:"bahan"`
	Satuan    string `json:"satuan"`
}

func FormatBahan(bahan entity.Bahan) BahanFormatter {
	bahanFormatter := BahanFormatter{}
	bahanFormatter.NamaBahan = bahan.NamaBahan
	bahanFormatter.Satuan = bahan.Satuan

	return bahanFormatter
}

func FormatBahans(datas []entity.Bahan) []BahanFormatter {
	bahanFormatters := []BahanFormatter{}
	for _, data := range datas {
		bahanFormatter := FormatBahan(data)
		bahanFormatters = append(bahanFormatters, bahanFormatter)
	}

	return bahanFormatters
}
