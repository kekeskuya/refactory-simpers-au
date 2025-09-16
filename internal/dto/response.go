package dto

type GetPersonelByNRPResponse struct {
	PersonelID   int    `json:"personel_id" example:"138"`
	PersonelNama string `json:"personel_nama" example:"Max Verstappen"`
	// Pangkat           string `json:"pangkat" example:"Marsekal Pertama TNI"`
	// Korps             string `json:"korps" example:"Lek"`
	// TempatLahir       string `json:"tempat_lahir" example:"Magelang"`
	// TanggalLahir      string `json:"tanggal_lahir" example:"21-05-1969"`
	// Jabatan           string `json:"jabatan" example:"Staff Khusus Kasau"`
	// Profesi           string `json:"profesi" example:"Elektronik"`
	// Spesialisasi      string `json:"spesialisasi" example:"Simulator"`
	// SatuanKerja       string `json:"satuan_kerja" example:"Kasau"`
	// TmtMasuk          string `json:"terhitung_masuk_tanggal" example:"28-09-1991"`
	// TmtPerwira        string `json:"terhitung_mulai_tanggal" example:"30-01-1997"`
	StatusPersonel_Id string `json:"StatusPersonel_Id" example:"1"`
}

type GetNPWPByNRPResponse struct {
	PersonelID int    `json:"personel_id" example:"138"`
	NPWP       string `json:"npwp" example:"916143918988772"`
}

type GetAsabriByNRPResponse struct {
	PersonelID int    `json:"personel_id" example:"138"`
	Asabri     string `json:"asabri" example:"AS4826315003"`
}

type GetPasporByNRPResponse struct {
	PersonelID  int    `json:"personel_id" example:"138"`
	NomorPaspor string `json:"nomor_paspor" example:"P27418523ID"`
}

type CreateLampiranResponse struct {
	PersonelID int `json:"personel_id" example:"1"`
}
