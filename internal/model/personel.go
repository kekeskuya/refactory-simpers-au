package model

import "database/sql"

type Personel struct {
	PersonelID              uint64         `db:"personel_id" json:"personel_id"`
	PersonelNama            string         `db:"personel_nama" json:"personel_nama"`
	PangkatNama             sql.NullString `db:"pangkat_nama" json:"pangkat_nama"`
	PangkatID               sql.NullInt32  `db:"pangkat_id" json:"pangkat_id"`
	KorpsNama               sql.NullString `db:"korps_nama" json:"korps_nama"`
	KorpsID                 sql.NullInt32  `db:"korps_id" json:"korps_id"`
	ProfesiID               sql.NullInt32  `db:"profesi_id" json:"profesi_id"`
	ProfesiNama             sql.NullString `db:"profesi_nama" json:"profesi_nama"`
	JabatanNamaPanjang      sql.NullString `db:"jabatan_nama_panjang" json:"jabatan_nama_panjang"`
	StatusPersonelID        sql.NullString `db:"statuspersonel_id" json:"statuspersonel_id"`
	NIK                     sql.NullString `db:"nik" json:"nik"`
	PendidikanAsalMasukID   sql.NullString `db:"pendidikan_asalmasuk_id" json:"pendidikan_asalmasuk_id"`
	PendidikanAsalMasukNama sql.NullString `db:"pendidikan_asalmasuk_nama" json:"pendidikan_asalmasuk_nama"`
	PendidikanMiliterID     sql.NullString `db:"pendidikan_militer_id" json:"pendidikan_militer_id"`
	PendidikanMiliterNama   sql.NullString `db:"pendidikan_militer_nama" json:"pendidikan_militer_nama"`
	PendidikanUmumID        sql.NullString `db:"pendidikan_umum_id" json:"pendidikan_umum_id"`
	PendidikanUmumNama      sql.NullString `db:"pendidikan_umum_nama" json:"pendidikan_umum_nama"`
	SatuanKerjaID           sql.NullString `db:"satuan_kerja_id" json:"satuan_kerja_id"`
	SatuanKerjaNama         sql.NullString `db:"satuan_kerja_nama" json:"satuan_kerja_nama"`
	// PersonelTanggalLahir    *time.Time     `db:"personel_tanggal_lahir" json:"personel_tanggal_lahir"`
}

func (p Personel) IsEmpty() bool {
	return p == Personel{}
}
