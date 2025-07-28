package model

import "time"

type Asabri struct {
	ID          uint64     `db:"id"`
	CreatedAt   *time.Time `db:"created_at"`
	UpdatedAt   *time.Time `db:"updated_at"`
	PersonelID  int        `db:"personel_id"`
	NomorAsabri string     `db:"nomor_asabri"`
}

func (a Asabri) IsEmpty() bool {
	return a == Asabri{}
}
