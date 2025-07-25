package model

import "time"

type Personel struct {
	ID              uint64     `db:"id"`
	CreatedAt       *time.Time `db:"created_at"`
	UpdatedAt       *time.Time `db:"updated_at"`
	Nama            string     `db:"nama"`
	NRP             string     `db:"nrp"`
	TmtMasuk        *time.Time `db:"tmt_masuk"`
	TmtPerwira      *time.Time `db:"tmt_perwira"`
	Pangkat         string     `db:"pangkat"`
	Korps           string     `db:"korps"`
	Profesi         string     `db:"profesi"`
	Spesialisasi    string     `db:"spesialisasi"`
	TempatLahir     string     `db:"tempat_lahir"`
	TanggalLahir    *time.Time `db:"tanggal_lahir"`
	Kesatuan        string     `db:"kesatuan"`
	Jabatan         string     `db:"jabatan"`
	StatusKeaktifan bool       `db:"status_keaktifan"`
}

func (p Personel) IsEmpty() bool {
	return p == Personel{}
}
