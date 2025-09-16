package entity

type AsabriWithNRP struct {
	PersonelID  int    `db:"personel_id"`
	NomorAsabri string `db:"asabri"`
}

func (a AsabriWithNRP) IsEmpty() bool {
	return a == AsabriWithNRP{}
}
