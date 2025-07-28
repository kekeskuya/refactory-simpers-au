package model

import "time"

type Paspor struct {
	ID            uint64     `db:"id"`
	CreatedAt     *time.Time `db:"created_at"`
	UpdatedAt     *time.Time `db:"updated_at"`
	PersonelID    int        `db:"personel_id"`
	NomorPassport string     `db:"nomor_passport"`
}

func (p Paspor) IsEmpty() bool {
	return p == Paspor{}
}
