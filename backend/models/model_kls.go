package models

type Kelas struct {
	ID       int    `json:"id"`
	Nama     string `json:"nama"`
	Kode     string `json:"kodeKelas"`
	nikDosen string `json:"nik"`
}
