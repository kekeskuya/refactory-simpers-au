package dto

type GetDikUmByNRPResponse struct {
	PersonelID            uint64 `json:"personel_id"`
	RPendidikanId         uint64 `json:"R_pendidikan_id"`
	JenisPendidikanNama   string `json:"jenis_pendidikan_nama"`
	SekolahNama           string `json:"sekolah_nama"`
	LembagaPendidikanNama string `json:"lembaga_pendidikan_nama"`
	SuratKeputusan        string `json:"surat_keputusan"`
	SuratKeputusanTgl     string `json:"surat_keputusan_tgl"`
	PendidikanTMT         string `json:"pendidikan_tmt"`
	TahunMasuk            string `json:"tahun_masuk"`
	TahunLulus            string `json:"tahun_lulus"`
	Gelar                 string `json:"gelar"`
}
