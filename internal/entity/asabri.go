package entity

import "time"

type AsabriWithNRP struct {
	ID          uint64     `db:"id"`
	CreatedAt   *time.Time `db:"created_at"`
	UpdatedAt   *time.Time `db:"updated_at"`
	PersonelID  int        `db:"personel_id"`
	NomorAsabri string     `db:"nomor_asabri"`
	NRP         string     `db:"nrp"`
}

func (a AsabriWithNRP) IsEmpty() bool {
	return a == AsabriWithNRP{}
}
