package dto

type GetJabatanByNRPResponse struct {
	PersonelID         uint64  `json:"personel_id,omitempty"`
	RJabatanId         string  `json:"r_jabatan_id,omitempty"`
	JabatanId          string  `json:"jabatan_id,omitempty"`
	JabatanNama        *string `json:"jabatan_nama,omitempty"`
	JabatanNamaPanjang *string `json:"jabatan_nama_panjang,omitempty"`
	SuratKeputusan     string  `json:"suratkeputusan,omitempty"`
	SuratKeputusanTgl  string  `json:"suratkeputusan_tgl,omitempty"`
	JabatanTMT         string  `json:"jabatan_tmt,omitempty"`
	JabatanTipe        string  `json:"jabatan_tipe,omitempty"`
	JabatanKeterangan  string  `json:"jabatan_keterangan,omitempty"`
	JabatanStatus      string  `json:"jabatan_status,omitempty"`
	SprintlakTgl       string  `json:"sprintlak_tgl,omitempty"`
}
