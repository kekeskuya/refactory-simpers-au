package model

import "database/sql"

type Pangkat struct {
	PersonelID          uint64         `db:"personel_id" json:"personel_id"`
	RPangkatId          sql.NullString `db:"r_pangkat_id" json:"r_pangkat_id"`
	PangkatId           sql.NullString `db:"pangkat_id" json:"pangkat_id"`
	PangkatNama         sql.NullString `db:"pangkat_nama" json:"pangkat_nama"`
	SuratKeputusan      sql.NullString `db:"surat_keputusan" json:"surat_keputusan"`
	SuratKeputusanTgl   sql.NullString `db:"surat_keputusan_tgl" json:"surat_keputusan_tgl"`
	SuratKeputusanJenis sql.NullString `db:"surat_keputusan_jenis" json:"surat_keputusan_jenis"`
	PangkatTMT          sql.NullString `db:"pangkat_tmt" json:"pangkat_tmt"`
	SprintlakTgl        sql.NullString `db:"sprintlak_tgl" json:"sprintlak_tgl"`
	SprintlakNo         sql.NullString `db:"sprintlak_no" json:"sprintlak_no"`
}
