package model

import "database/sql"

type DikUm struct {
	PersonelID            uint64         `db:"personel_id" json:"personel_id"`
	RPendidikanId         uint64         `db:"R_pendidikan_id" json:"R_pendidikan_id"`
	JenisPendidikanNama   sql.NullString `db:"JenisPendidikan_Nama" json:"JenisPendidikan_Nama"`
	SekolahNama           sql.NullString `db:"Sekolah_Nama" json:"Sekolah_Nama"`
	LembagaPendidikanNama sql.NullString `db:"LembagaPendidikan_Nama" json:"LembagaPendidikan_Nama"`
	SuratKeputusan        sql.NullString `db:"SuratKeputusan" json:"SuratKeputusan"`
	SuratKeputusanTgl     sql.NullString `db:"SuratKeputusan_Tgl" json:"SuratKeputusan_Tgl"`
	PendidikanTMT         sql.NullString `db:"Pendidikan_TMT" json:"Pendidikan_TMT"`
	TahunMasuk            sql.NullString `db:"Tahun_Masuk" json:"Tahun_Masuk"`
	TahunLulus            sql.NullString `db:"Tahun_Lulus" json:"Tahun_Lulus"`
	Gelar                 sql.NullString `db:"Gelar" json:"Gelar"`
}

// func (p FamilyCard) IsEmpty() bool {
// 	return p == FamilyCard{}
// }
