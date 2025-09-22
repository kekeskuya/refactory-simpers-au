package dto

type GetLampiranKKByNRPResponse struct {
	PersonelID         uint64 `json:"personel_id,omitempty"`
	DKartuKeluargaID   string `json:"d_kartu_keluarga_id,omitempty"`
	DKartuKeluargaNama string `json:"d_kartu_keluarga_nama,omitempty"`
}
