package model

import (
	"database/sql"
	"time"
)

type DokumenLampiran struct {
	ID         uint64         `db:"id" json:"id"`
	PersonelID uint64         `db:"personel_id" json:"personel_id"`
	RID        sql.NullString `db:"r_id" json:"r_id"`
	RTipe      sql.NullString `db:"r_tipe" json:"r_tipe"`
	CreateDate time.Time      `db:"create_date" json:"create_date"`
	URL        sql.NullString `db:"url" json:"url"`
}
