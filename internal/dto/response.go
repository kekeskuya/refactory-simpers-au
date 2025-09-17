package dto

import (
	"refactory-simpers-au/internal/model"
)

type GetPersonelByNRPResponse struct {
	PersonelID              uint64  `json:"personel_id"`
	PersonelNama            string  `json:"personel_nama"`
	PangkatNama             string  `json:"pangkat_nama,omitempty"`
	PangkatID               *int32  `json:"pangkat_id,omitempty"`
	KorpsNama               string  `json:"korps_nama,omitempty"`
	KorpsID                 *int32  `json:"korps_id,omitempty"`
	ProfesiID               *int32  `json:"profesi_id,omitempty"`
	ProfesiNama             string  `json:"profesi_nama,omitempty"`
	StatusPersonelID        *string `json:"statuspersonelid,omitempty"`
	JabatanNamaPanjang      string  `json:"jabatan_nama_panjang,omitempty"`
	NIK                     string  `json:"nik,omitempty"`
	PendidikanAsalMasukID   string  `json:"pendidikan_asalmasuk_id,omitempty"`
	PendidikanAsalMasukNama string  `json:"pendidikan_asalmasuk_nama,omitempty"`
	PendidikanMiliterID     string  `json:"pendidikan_militer_id,omitempty"`
	PendidikanMiliterNama   string  `json:"pendidikan_militer_nama,omitempty"`
	PendidikanUmumID        string  `json:"pendidikan_umum_id,omitempty"`
	PendidikanUmumNama      string  `json:"pendidikan_umum_nama,omitempty"`
	SatuanKerjaID           string  `json:"satuan_kerja_id,omitempty"`
	SatuanKerjaNama         string  `json:"satuan_kerja_nama,omitempty"`
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
	} else {
		profesiID = nil
	}

	var statusPersonelID *string
	if m.StatusPersonelID.Valid {
		statusPersonelID = &m.StatusPersonelID.String
	} else {
		statusPersonelID = nil
	}

	var NIK string
	if m.NIK.Valid {
		NIK = m.NIK.String
	} else {
		NIK = "XXX"
	}

	var JabatanNamaPanjang string
	if m.JabatanNamaPanjang.Valid {
		JabatanNamaPanjang = m.JabatanNamaPanjang.String
	} else {
		JabatanNamaPanjang = "-"
	}

	var pendidikanAsalMasukID string
	if m.PendidikanAsalMasukID.Valid {
		pendidikanAsalMasukID = m.PendidikanAsalMasukID.String
	} else {
		pendidikanAsalMasukID = "-"
	}

	var pendidikanAsalMasukNama string
	if m.PendidikanAsalMasukNama.Valid {
		pendidikanAsalMasukNama = m.PendidikanAsalMasukNama.String
	} else {
		pendidikanAsalMasukNama = "-"
	}

	var pendidikanMiliterID string
	if m.PendidikanMiliterID.Valid {
		pendidikanMiliterID = m.PendidikanMiliterID.String
	} else {
		pendidikanMiliterID = "-"
	}

	var pendidikanMiliterNama string
	if m.PendidikanMiliterNama.Valid {
		pendidikanMiliterNama = m.PendidikanMiliterNama.String
	} else {
		pendidikanMiliterNama = "-"
	}

	var pendidikanUmumID string
	if m.PendidikanUmumID.Valid {
		pendidikanUmumID = m.PendidikanUmumID.String
	} else {
		pendidikanUmumID = "-"
	}

	var pendidikanUmumNama string
	if m.PendidikanUmumNama.Valid {
		pendidikanUmumNama = m.PendidikanUmumNama.String
	} else {
		pendidikanUmumNama = "-"
	}

	var satuan_kerja_id string
	if m.SatuanKerjaID.Valid {
		satuan_kerja_id = m.SatuanKerjaID.String
	} else {
		satuan_kerja_id = "-"
	}

	var satuan_kerja_nama string
	if m.SatuanKerjaNama.Valid {
		satuan_kerja_nama = m.SatuanKerjaNama.String
	} else {
		satuan_kerja_nama = "-"
	}

	return GetPersonelByNRPResponse{
		PersonelID:              m.PersonelID,
		PersonelNama:            m.PersonelNama,
		PangkatNama:             pangkatNama,
		PangkatID:               pangkatID,
		KorpsNama:               korpsNama,
		KorpsID:                 korpsID,
		ProfesiID:               profesiID,
		ProfesiNama:             profesiNama,
		StatusPersonelID:        statusPersonelID,
		NIK:                     NIK,
		JabatanNamaPanjang:      JabatanNamaPanjang,
		PendidikanAsalMasukID:   pendidikanAsalMasukID,
		PendidikanAsalMasukNama: pendidikanAsalMasukNama,
		PendidikanMiliterID:     pendidikanMiliterID,
		PendidikanMiliterNama:   pendidikanMiliterNama,
		PendidikanUmumID:        pendidikanUmumID,
		PendidikanUmumNama:      pendidikanUmumNama,
		SatuanKerjaID:           satuan_kerja_id,
		SatuanKerjaNama:         satuan_kerja_nama,
		// PersonelTanggalLahir: tmtMasuk,
	}
}
