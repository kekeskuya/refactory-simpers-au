package model

import "database/sql"

type Jabatan struct {
	PersonelID         uint64         `db:"personel_id" json:"personel_id"`
	RJabatanId         sql.NullString `db:"r_jabatan_id" json:"r_jabatan_id"`
	JabatanId          sql.NullString `db:"jabatan_id" json:"jabatan_id"`
	JabatanNama        sql.NullString `db:"jabatan_nama" json:"jabatan_nama"`
	JabatanNamaPanjang sql.NullString `db:"jabatan_nama_panjang" json:"jabatan_nama_panjang"`
	SuratKeputusan     sql.NullString `db:"suratkeputusan" json:"suratkeputusan"`
	SuratKeputusanTgl  sql.NullString `db:"suratkeputusan_tgl" json:"suratkeputusan_tgl"`
	JabatanTMT         sql.NullString `db:"jabatan_tmt" json:"jabatan_tmt"`
	JabatanTipe        sql.NullString `db:"r_jabatan_tipe" json:"r_jabatan_tipe"`
	JabatanKeterangan  sql.NullString `db:"r_jabatan_keterangan" json:"r_jabatan_keterangan"`
	JabatanStatus      sql.NullString `db:"jabatan_status" json:"jabatan_status"`
	SprintlakTgl       sql.NullString `db:"sprintlak_tgl" json:"sprintlak_tgl"`
}
