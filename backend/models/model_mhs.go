package models

type (
	Mahasiswa struct {
		Nama     string `json:"nama"`
		Nim      string `json:"nim"`
		Password string `json:"pass"`
	}

	ResponseMahasiswa struct {
		Status  int         `json:"status"`
		Message string      `json:"message"`
		Data    []Mahasiswa `json:"data"`
	}

	MhsArray []Mahasiswa
)
