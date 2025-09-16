package model

type Paspor struct {
	PersonelID    int    `db:"personel_id"`
	NomorPassport string `db:"nomor_passport"`
}

func (p Paspor) IsEmpty() bool {
	return p == Paspor{}
}
