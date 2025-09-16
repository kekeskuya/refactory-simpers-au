package model

type NPWP struct {
	PersonelID int    `db:"personel_id"`
	NPWP       string `db:"npwp"`
}

func (n NPWP) IsEmpty() bool {
	return n == NPWP{}
}
