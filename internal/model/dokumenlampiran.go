package model

import (
	"database/sql"
	"time"
)

type DokumenLampiran struct {
	ID         sql.NullInt64  `db:"id" json:"id"`
	PersonelID sql.NullInt64  `db:"personel_id" json:"personel_id"`
	RID        sql.NullString `db:"r_id" json:"r_id"`
	RTipe      sql.NullString `db:"r_tipe" json:"r_tipe"`
	CreateDate time.Time      `db:"create_date" json:"create_date"`
	URL        sql.NullString `db:"url" json:"url"`
}

// type DokumenLampiran struct {
// 	RID   string `json:"r_id"`
// 	RTipe string `json:"r_tipe"`
// 	URL   string `json:"url"`
// }
