package controllers

import (
	"college-with-ui/backend/database"
	"college-with-ui/backend/models"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

var nilai models.Nilai

func InputNilaiMhs(ctx echo.Context) error {
	db := database.Connect()
	defer db.Close()

	if err := ctx.Bind(nilai); err != nil {
		return err
	}

	total := totalNilai(nilai)

	query := "INSERT INTO nilai(nim_mhs, nik_dosen, kode_mk, absen, nilai, total_nilai) VALUES(?,?,?,?,?,?)"
	stmt, err := db.Prepare(query)

	if err != nil {
		fmt.Print(err.Error())
	}
	defer stmt.Close()
	result, err2 := stmt.Exec(nilai.NIM, nilai.NIK, nilai.Kode, nilai.Absen, nilai.Nilai, total)

	// Exit if we get an error
	if err2 != nil {
		panic(err2)
	}
	fmt.Println(result.LastInsertId())

	return ctx.JSON(http.StatusCreated, nilai)
}

func totalNilai(nilai models.Nilai) float64 {

	bobotAbsen := (float64(nilai.Absen) / 14) * 30
	bobotNilai := nilai.Nilai * 70 / 100
	totalNilai := bobotAbsen + bobotNilai

	fmt.Println(bobotAbsen)
	fmt.Println(bobotNilai)
	fmt.Println(totalNilai)
	return totalNilai
}
