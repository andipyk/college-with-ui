package models

type (
	Admin struct {
		Username string `json:"user"`
		Password string `json:"pass"`
	}

	ResponseAdmin struct {
		Status  int     `json:"status"`
		Message string  `json:"message"`
		Data    []Admin `json:"data"`
	}

	AdmArray []Admin
)
