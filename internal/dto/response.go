package dto

import (
	"refactory-simpers-au/internal/model"
)

type GetPersonelByNRPResponse struct {
	PersonelID         uint64  `json:"personel_id"`
	PersonelNama       string  `json:"personel_nama"`
	PangkatNama        string  `json:"pangkat_nama,omitempty"`
	PangkatID          *int32  `json:"pangkat_id,omitempty"`
	KorpsNama          string  `json:"korps_nama,omitempty"`
	KorpsID            *int32  `json:"korps_id,omitempty"`
	ProfesiID          *int32  `json:"profesi_id,omitempty"`
	ProfesiNama        string  `json:"profesi_nama,omitempty"`
	StatusPersonelID   *string `json:"statuspersonelid,omitempty"`
	JabatanNamaPanjang string  `json:"jabatan_nama_panjang,omitempty"`
	// PersonelTanggalLahir string  `json:"personel_tanggal_lahir,omitempty"`
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

func ToGetPersonelByNRPResponse(m model.Personel) GetPersonelByNRPResponse {
	var pangkatNama string
	if m.PangkatNama.Valid {
		pangkatNama = m.PangkatNama.String
	} else {
		pangkatNama = "-"
	}

	var pangkatID *int32
	if m.PangkatID.Valid {
		pangkatID = &m.PangkatID.Int32
	}

	var korpsNama string
	if m.KorpsNama.Valid {
		korpsNama = m.KorpsNama.String
	} else {
		korpsNama = "-"
	}

	var korpsID *int32
	if m.KorpsID.Valid {
		korpsID = &m.KorpsID.Int32
	}

	var profesiNama string
	if m.ProfesiNama.Valid {
		profesiNama = m.ProfesiNama.String
	} else {
		profesiNama = "-"
	}

	var profesiID *int32
	if m.ProfesiID.Valid {
		profesiID = &m.ProfesiID.Int32
	}

	var statusPersonelID *string
	if m.StatusPersonelID.Valid {
		statusPersonelID = &m.StatusPersonelID.String
	}

	return GetPersonelByNRPResponse{
		PersonelID:       m.PersonelID,
		PersonelNama:     m.PersonelNama,
		PangkatNama:      pangkatNama,
		PangkatID:        pangkatID,
		KorpsNama:        korpsNama,
		KorpsID:          korpsID,
		ProfesiID:        profesiID,
		ProfesiNama:      profesiNama,
		StatusPersonelID: statusPersonelID,
	}
}
