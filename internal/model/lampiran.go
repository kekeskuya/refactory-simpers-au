package model

import (
	"refactory-simpers-au/constants"
	"time"
)

type Lampiran struct {
	ID         uint64             `db:"id"`
	CreatedAt  *time.Time         `db:"created_at"`
	UpdatedAt  *time.Time         `db:"updated_at"`
	Kategori   constants.Category `db:"kategori"`
	PersonelID int                `db:"personel_id"`
	DokumenID  string             `db:"dokumen_id"`
	Link       string             `db:"link"`
	Nama       string             `db:"nama"`
	Keterangan string             `db:"keterangan"`
	Tipe       string             `db:"tipe"`
}

func (l Lampiran) IsEmpty() bool {
	return l == Lampiran{}
}
