package entity

type PasporWithNRP struct {
	PersonelID    int    `db:"personel_id"`
	NomorPassport string `db:"passport_id"`
}

func (p PasporWithNRP) IsEmpty() bool {
	return p == PasporWithNRP{}
}
