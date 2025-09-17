package model

import "database/sql"

type FamilyCard struct {
	PersonelID           uint64         `db:"personel_id" json:"personel_id"`
	NamaKeluarga         sql.NullString `db:"nama_keluarga" json:"nama_keluarga"`
	AlamatKeluarga       sql.NullString `db:"alamat_keluarga" json:"alamat_keluarga"`
	TempatLahirKeluarga  sql.NullString `db:"tempat_lahir_keluarga" json:"tempat_lahir_keluarga"`
	TanggalLahirKeluarga sql.NullString `db:"tanggal_lahir_keluarga" json:"tanggal_lahir_keluarga"`
	JenisKelaminKeluarga sql.NullString `db:"jenis_kelamin_keluarga" json:"jenis_kelamin_keluarga"`
	StatusNikahKeluarga  sql.NullString `db:"status_nikah_keluarga" json:"status_nikah_keluarga"`
	PekerjaanKeluarga    sql.NullString `db:"pekerjaan_keluarga" json:"pekerjaan_keluarga"`
	HubunganKeluarga     sql.NullString `db:"hubungan_keluarga" json:"hubungan_keluarga"`
}

// func (p FamilyCard) IsEmpty() bool {
// 	return p == FamilyCard{}
// }
