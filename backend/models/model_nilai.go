package models

type (
	Nilai struct {
		ID         int     `json:"id"`
		NIM        string  `json:"nim"`
		NIKDosen   string  `json:"nik"`
		KodeMK     string  `json:"kode"`
		Absen      int64   `json:"absen"`
		Nilai      float64 `json:"nilai"`
		TotalNilai float64 `json:"total_nilai"`
	}

	ResponseReport struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
		Data    Nilai  `json:"data"`
	}

	nilaiArr []Nilai
)
