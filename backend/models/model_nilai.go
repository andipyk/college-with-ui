package models

type (
	Nilai struct {
		NIM        string  `json:"nim"`
		NamaMhs    string  `json:"namamhs"`
		NIKDosen   string  `json:"nik"`
		KodeMK     string  `json:"kode"`
		NamaMK     string  `json:"namamk"`
		Absen      int64   `json:"absen"`
		Nilai      float64 `json:"nilai"`
		TotalNilai float64 `json:"total_nilai"`
	}

	ResponseReport struct {
		Status  int     `json:"status"`
		Message string  `json:"message"`
		Data    []Nilai `json:"data"`
	}

	ReportArray []Nilai
)
