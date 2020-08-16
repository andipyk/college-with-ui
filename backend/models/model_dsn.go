package models

type (
	Dosen struct {
		ID   int    `json:"id"`
		Nama string `json:"nama"`
		NIK  string `json:"nik"`
	}

	ResponseDosen struct {
		Status  int     `json:"status"`
		Message string  `json:"message"`
		Data    []Dosen `json:"data"`
	}

	DsnArray []Dosen
)
