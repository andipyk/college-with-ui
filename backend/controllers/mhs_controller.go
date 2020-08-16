package controllers

import (
	"college-with-ui/backend/database"
	"college-with-ui/backend/models"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

var response models.ResponseMahasiswa

func GetAllMahasiswa(ctx echo.Context) error {
	var arrMhs []models.Mahasiswa
	var mhs models.Mahasiswa

	db := database.Connect()
	defer db.Close()

	rows, err := db.Query("SELECT nama, nim FROM mahasiswa")
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&mhs.Nama, &mhs.Nim); err != nil {
			log.Fatal(err.Error())
		} else {
			arrMhs = append(arrMhs, mhs)
		}
	}

	response.Status = 1
	response.Message = "Success"
	response.Data = arrMhs

	// reset kembali splice
	arrMhs = arrMhs[:0]
	return ctx.JSON(http.StatusOK, response)
}

func CreateMahasiswa(ctx echo.Context) error {
	var arrMhs []models.Mahasiswa

	db := database.Connect()
	defer db.Close()

	mhs := &models.Mahasiswa{}

	if err := ctx.Bind(mhs); err != nil {
		return err
	}

	_, err := db.Exec("INSERT INTO mahasiswa (nama, nim) values (?,?)",
		mhs.Nama,
		mhs.Nim,
	)

	if err != nil {
		log.Fatal(err.Error())
	}

	arrMhs = append(arrMhs, models.Mahasiswa{Nim: mhs.Nim, Nama: mhs.Nama})

	response.Status = 1
	response.Message = "Success Add"
	response.Data = arrMhs

	log.Print("Insert data to database")

	return ctx.JSON(http.StatusCreated, response)
}
