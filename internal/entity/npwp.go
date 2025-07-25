package entity

import "time"

type NPWPWithNRP struct {
	ID         uint64     `db:"id"`
	CreatedAt  *time.Time `db:"created_at"`
	UpdatedAt  *time.Time `db:"updated_at"`
	PersonelID int        `db:"personel_id"`
	NPWP       string     `db:"npwp"`
	NRP        string     `db:"nrp"`
}

func (n NPWPWithNRP) IsEmpty() bool {
	return n == NPWPWithNRP{}
}
