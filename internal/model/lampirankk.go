package model

import "database/sql"

type LampiranKK struct {
	PersonelID         uint64         `db:"personel_id" json:"personel_id"`
	DKartuKeluargaID   sql.NullString `db:"d_kartu_keluarga_id" json:"d_kartu_keluarga_id"`
	DKartuKeluargaNama sql.NullString `db:"d_kartu_keluarga_nama" json:"d_kartu_keluarga_nama"`
}
