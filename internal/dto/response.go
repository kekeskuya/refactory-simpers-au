package dto

type GetPersonelByNRPResponse struct {
	ID              int    `json:"id" example:"138"`
	NRP             string `json:"nrp" example:"120398109238"`
	Nama            string `json:"nama" example:"Max Verstappen"`
	Pangkat         string `json:"pangkat" example:"Marsekal Pertama TNI"`
	Korps           string `json:"korps" example:"Lek"`
	StatusKeaktifan bool   `json:"status_keaktifan" example:"true"`
	TempatLahir     string `json:"tempat_lahir" example:"Magelang"`
	TanggalLahir    string `json:"tanggal_lahir" example:"21-05-1969"`
	Jabatan         string `json:"jabatan" example:"Staff Khusus Kasau"`
	Profesi         string `json:"profesi" example:"Elektronik"`
	Spesialisasi    string `json:"spesialisasi" example:"Simulator"`
	SatuanKerja     string `json:"satuan_kerja" example:"Kasau"`
	TmtMasuk        string `json:"terhitung_masuk_tanggal" example:"28-09-1991"`
	TmtPerwira      string `json:"terhitung_mulai_tanggal" example:"30-01-1997"`
}

type GetNPWPByNRPResponse struct {
	ID   int    `json:"id" example:"138"`
	NRP  string `json:"nrp" example:"120398109238"`
	NPWP string `json:"npwp" example:"916143918988772"`
}

type GetAsabriByNRPResponse struct {
	ID     int    `json:"id" example:"138"`
	NRP    string `json:"nrp" example:"120398109238"`
	Asabri string `json:"asabri" example:"AS4826315003"`
}

type GetPasporByNRPResponse struct {
	ID          int    `json:"id" example:"138"`
	NRP         string `json:"nrp" example:"120398109238"`
	NomorPaspor string `json:"nomor_paspor" example:"P27418523ID"`
}

type CreateLampiranResponse struct {
	ID int `json:"id" example:"1"`
}
