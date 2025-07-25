package model

import "time"

type NPWP struct {
	ID         uint64     `db:"id"`
	CreatedAt  *time.Time `db:"created_at"`
	UpdatedAt  *time.Time `db:"updated_at"`
	PersonelID int        `db:"personel_id"`
	NPWP       string     `db:"npwp"`
}

func (n NPWP) IsEmpty() bool {
	return n == NPWP{}
}
