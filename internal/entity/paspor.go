package entity

import "time"

type PasporWithNRP struct {
	ID            uint64     `db:"id"`
	CreatedAt     *time.Time `db:"created_at"`
	UpdatedAt     *time.Time `db:"updated_at"`
	PersonelID    int        `db:"personel_id"`
	NomorPassport string     `db:"nomor_passport"`
	NRP           string     `db:"nrp"`
}

func (p PasporWithNRP) IsEmpty() bool {
	return p == PasporWithNRP{}
}
