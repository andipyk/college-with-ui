package models

type (
	Mahasiswa struct {
		Nama string `json:"nama"`
		Nim  string `json:"nim"`
	}

	ResponseMahasiswa struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
		Data    []Mahasiswa
	}
)
