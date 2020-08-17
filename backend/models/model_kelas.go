package models

type (
	Kelas struct {
		ID       int    `json:"id"`
		Nama     string `json:"nama"`
		KodeMK   string `json:"kode"`
		NIKDosen string `json:"nik"`
	}

	ResponseKelas struct {
		Status  int     `json:"status"`
		Message string  `json:"message"`
		Data    []Kelas `json:"data"`
	}

	KlsArray []Kelas
)
