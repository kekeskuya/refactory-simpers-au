package dto

type GetPangkatByNRPResponse struct {
	PersonelID          uint64  `json:"personel_id,omitempty"`
	RPangkatId          string  `json:"r_pangkat_id,omitempty"`
	PangkatId           string  `json:"pangkat_id,omitempty"`
	PangkatNama         *string `json:"pangkat_nama,omitempty"`
	SuratKeputusan      string  `json:"suratkeputusan,omitempty"`
	SuratKeputusanTgl   string  `json:"suratkeputusan_tgl,omitempty"`
	SuratKeputusanJenis string  `json:"surat_keputusan_jenis,omitempty"`
	PangkatTMT          string  `json:"pangkat_tmt,omitempty"`
	SprintlakTgl        string  `json:"sprintlak_tgl,omitempty"`
	SprintlakNo         string  `json:"sprintlak_no,omitempty"`
}
