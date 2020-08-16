package models

type (
	Nilai struct {
		ID         int     `json:"id"`
		NIM        string  `json:"nim"`
		NIK        string  `json:"nik"`
		Kode       string  `json:"kode"`
		Absen      int64   `json:"absen"`
		Nilai      float64 `json:"nilai"`
		TotalNilai float64 `json:"total_nilai"`
	}
)
