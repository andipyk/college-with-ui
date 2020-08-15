package models

type (
	Nilai struct {
		ID         int     `json:"id"`
		nimMhs     string  `json:"nim"`
		nikDosen   string  `json:"nik"`
		kodeKelas  string  `json:"kodeKelas"`
		Absen      int     `json:"absen"`
		Nilai      float32 `json:"nilai"`
		totalNilai float32 `json:"tNilai"`
	}
)
