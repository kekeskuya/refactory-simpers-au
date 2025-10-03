package dto

import "time"

// type DokumenLampiranByNRPRequest struct {
// 	LampiranID int       `json:"lampiran_id"`
// 	PersonelID uint64    `json:"personel_id"`
// 	RID        string    `json:"r_id"`
// 	RTipe      string    `json:"r_tipe"`
// 	CreateDate time.Time `json:"create_date"`
// 	LinkUrl    string    `json:"url"`
// }

// type PostDokumenLampiranByNRPResponse struct {
// 	LampiranID int       `json:"lampiran_id"`
// 	PersonelID uint64    `json:"personel_id"`
// 	RID        string    `json:"r_id"`
// 	RTipe      string    `json:"r_tipe"`
// 	CreateDate time.Time `json:"create_date"`
// 	LinkUrl    string    `json:"url"`
// }

type DokumenLampiranByNRPRequest struct {
	ID         int       `json:"id"`
	PersonelID int64     `json:"personel_id"`
	RID        string    `json:"r_id"`
	RTipe      string    `json:"r_tipe"`
	CreateDate time.Time `json:"create_date"`
	URL        string    `json:"url"`
}

type PostDokumenLampiranByNRPResponse struct {
	ID         int       `json:"id"`
	PersonelID int64     `json:"personel_id"`
	RID        string    `json:"r_id"`
	RTipe      string    `json:"r_tipe"`
	CreateDate time.Time `json:"create_date"`
	URL        string    `json:"url"`
}
