package controllers

import (
	"college-with-ui/backend/database"
	"college-with-ui/backend/models"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

var arrKelas models.KlsArray
var resKls models.ResponseKelas

func AddKelas(ctx echo.Context) error {
	db := database.Connect()
	defer db.Close()

	// bind struct dengan message body
	kls := &models.Kelas{}
	if err := ctx.Bind(kls); err != nil {
		return err
	}

	// buat query untuk DB
	query := "INSERT INTO kelas(nama, kode, nik_dosen) VALUES (?,?,?)"
	stmt, err := db.Prepare(query)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer stmt.Close()

	// untuk print di midware
	result, err2 := stmt.Exec(kls.Nama, kls.KodeMK, kls.NIKDosen)
	// keluar apabila ada error
	if err2 != nil {
		panic(err2)
	}
	log.Println(result.LastInsertId())

	// untuk memberi respon ke client
	arrKelas = append(arrKelas, models.Kelas{Nama: kls.Nama, KodeMK: kls.KodeMK, NIKDosen: kls.NIKDosen})

	resKls.Status = 1
	resKls.Message = "Success Add Kelas"
	resKls.Data = arrKelas

	return ctx.JSON(http.StatusCreated, resKls)
}

func GetAllKelasByMK(ctx echo.Context) error {
	db := database.Connect()
	defer db.Close()

	var kls models.Kelas

	requestedMK := ctx.Param("kode")
	rows, err := db.Query("SELECT nama, kode, nik_dosen FROM kelas WHERE kode = ?", requestedMK)
	if err != nil {
		log.Fatal(err.Error())
	}

	for rows.Next() {
		if err := rows.Scan(&kls.Nama, &kls.KodeMK, &kls.NIKDosen); err != nil {
			log.Fatal(err.Error())
		} else {
			arrKelas = append(arrKelas, kls)
		}
	}

	resKls.Status = 1
	resKls.Message = "Success"
	resKls.Data = arrKelas

	// reset kembali splice
	arrDosen = arrDosen[:0]
	return ctx.JSON(http.StatusOK, resKls)
}
