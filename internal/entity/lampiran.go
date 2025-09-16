package entity

type LampiranWithNRP struct {
	PersonelID int    `db:"personel_id"`
	Nama       string `db:"nama"`
}

func (l LampiranWithNRP) IsEmpty() bool {
	return l == LampiranWithNRP{}
}
