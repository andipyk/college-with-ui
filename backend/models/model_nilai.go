package models

type (
	Nilai struct {
		ID         int     `json:"id"`
		NIM        string  `json:"nim"`
		NIK        string  `json:"nik"`
		Kode       string  `json:"kode"`
		Absen      int     `json:"absen"`
		Nilai      float32 `json:"nilai"`
		TotalNilai float32 `json:"total_nilai"`
	}
)
