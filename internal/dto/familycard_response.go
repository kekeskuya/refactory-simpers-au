package dto

type GetFamilyCardByNRPResponse struct {
	PersonelID           uint64 `json:"personel_id"`
	NamaKeluarga         string `json:"nama_keluarga"`
	AlamatKeluarga       string `json:"alamat_keluarga"`
	TempatLahirKeluarga  string `json:"tempat_lahir_keluarga"`
	TanggalLahirKeluarga string `json:"tanggal_lahir_keluarga"`
	JenisKelaminKeluarga string `json:"jenis_kelamin_keluarga"`
	StatusNikahKeluarga  string `json:"status_nikah_keluarga"`
	PekerjaanKeluarga    string `json:"pekerjaan_keluarga"`
	HubunganKeluarga     string `json:"hubungan_keluarga"`
}
