package model

type Asabri struct {
	PersonelID  int    `db:"personel_id"`
	NomorAsabri string `db:"nomor_asabri"`
}

func (a Asabri) IsEmpty() bool {
	return a == Asabri{}
}
