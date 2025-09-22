package dto

type GetTanhorByNRPResponse struct {
	PersonelID        uint64  `json:"personel_id,omitempty"`
	RJasaId           string  `json:"r_jasa_id,omitempty"`
	JasaId            string  `json:"jasa_id,omitempty"`
	JasaNama          *string `json:"jasa_nama,omitempty"`
	SuratKeputusan    string  `json:"suratkeputusan,omitempty"`
	SuratKeputusanTgl string  `json:"suratkeputusan_tgl,omitempty"`
	JasaTMT           string  `json:"jasa_tmt,omitempty"`
	JasaTST           string  `json:"jasa_tst,omitempty"`
	JasaTipe          string  `json:"jasa_tipe,omitempty"`
	Keterangan        string  `json:"keterangan,omitempty"`
}
