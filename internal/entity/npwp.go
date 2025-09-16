package entity

type NPWPWithNRP struct {
	PersonelID int    `db:"personel_id"`
	NPWP       string `db:"npwp"`
}

func (n NPWPWithNRP) IsEmpty() bool {
	return n == NPWPWithNRP{}
}
