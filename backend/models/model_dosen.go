package models

type (
	Dosen struct {
		Nama     string `json:"nama"`
		NIK      string `json:"nik"`
		Password string `json:"pass"`
	}

	ResponseDosen struct {
		Status  int     `json:"status"`
		Message string  `json:"message"`
		Data    []Dosen `json:"data"`
	}

	DosArray []Dosen
)
