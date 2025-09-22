package model

import "database/sql"

type Tanhor struct {
	PersonelID        uint64         `db:"personel_id" json:"personel_id"`
	RJasaId           sql.NullString `db:"r_jasa_id" json:"r_jasa_id"`
	JasaId            sql.NullString `db:"jasa_id" json:"jasa_id"`
	JasaNama          sql.NullString `db:"jasa_nama" json:"jasa_nama"`
	SuratKeputusan    sql.NullString `db:"suratkeputusan" json:"suratkeputusan"`
	SuratKeputusanTgl sql.NullString `db:"suratkeputusan_tgl" json:"suratkeputusan_tgl"`
	JasaTMT           sql.NullString `db:"jasa_tmt" json:"jasa_tmt"`
	JasaTST           sql.NullString `db:"jasa_tst" json:"jasa_tst"`
	Keterangan        sql.NullString `db:"keterangan" json:"keterangan"`
}
