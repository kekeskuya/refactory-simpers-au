package entity

import (
	"refactory-simpers-au/constants"
	"time"
)

type LampiranWithNRP struct {
	ID         uint64             `db:"id"`
	CreatedAt  *time.Time         `db:"created_at"`
	UpdatedAt  *time.Time         `db:"updated_at"`
	Kategori   constants.Category `db:"kategori"`
	PersonelID int                `db:"personel_id"`
	Link       string             `db:"link"`
	Nama       string             `db:"nama"`
	Keterangan string             `db:"keterangan"`
	Tipe       string             `db:"tipe"`
	NRP        string             `db:"nrp"`
}

func (l LampiranWithNRP) IsEmpty() bool {
	return l == LampiranWithNRP{}
}
