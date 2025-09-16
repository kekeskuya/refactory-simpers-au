package model

type Personel struct {
	PersonelID   uint64 `db:"personel_id"`
	PersonelNama string `db:"Personel_Nama"`
	// NRP               string     `db:"nrp"`
	// TmtMasuk          *time.Time `db:"tmt_masuk"`
	// TmtPerwira        *time.Time `db:"tmt_perwira"`
	// Pangkat           string     `db:"pangkat"`
	// Korps             string     `db:"korps"`
	// Profesi           string     `db:"profesi"`
	// Spesialisasi      string     `db:"spesialisasi"`
	// TempatLahir       string     `db:"tempat_lahir"`
	// TanggalLahir      *time.Time `db:"tanggal_lahir"`
	// Kesatuan          string     `db:"kesatuan"`
	// Jabatan           string     `db:"jabatan"`
	StatusPersonel_Id string `db:"StatusPersonel_Id"`
}

func (p Personel) IsEmpty() bool {
	return p == Personel{}
}
