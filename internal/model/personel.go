package model

import "database/sql"

type Personel struct {
	PersonelID         uint64         `db:"personel_id" json:"personel_id"`
	PersonelNama       string         `db:"personel_nama" json:"personel_nama"`
	PangkatNama        sql.NullString `db:"pangkat_nama" json:"pangkat_nama"`
	PangkatID          sql.NullInt32  `db:"pangkat_id" json:"pangkat_id"`
	KorpsNama          sql.NullString `db:"korps_nama" json:"korps_nama"`
	KorpsID            sql.NullInt32  `db:"korps_id" json:"korps_id"`
	ProfesiID          sql.NullInt32  `db:"profesi_id" json:"profesi_id"`
	ProfesiNama        sql.NullString `db:"profesi_nama" json:"profesi_nama"`
	JabatanNamaPanjang sql.NullString `db:"jabatan_nama_panjang" json:"jabatan_nama_panjang"`
	//PersonelTanggalLahir sql.NullString `db:"personel_tanggal_lahir" json:"personel_tanggal_lahir"`
	StatusPersonelID sql.NullString `db:"statuspersonel_id" json:"statuspersonel_id"`
}

func (p Personel) IsEmpty() bool {
	return p == Personel{}
}
