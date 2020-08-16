package controllers

import (
	"college-with-ui/backend/database"
	"college-with-ui/backend/models"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

var mhs models.Mahasiswa
var arrMhs []models.Mahasiswa
var response models.ResponseMahasiswa

func GetAllMahasiswa(ctx echo.Context) error {
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

func AddMahasiswa(ctx echo.Context) error {
	db := database.Connect()
	defer db.Close()

	if err := ctx.Bind(mhs); err != nil {
		return err
	}

	// query, err := db.Exec("INSERT INTO mahasiswa VALUES (?,?)")
	// query.Exec("nama, nim")
	query := "INSERT INTO mahasiswa(nama, nim) VALUES (?,?)"
	stmt, err := db.Prepare(query)
	if err != nil {
		fmt.Print(err.Error())
	}

	defer stmt.Close()
	result, err2 := stmt.Exec(mhs.Nama, mhs.Nim)

	// keluar apabila ada error
	if err2 != nil {
		panic(err2)
	}
	fmt.Println(result.LastInsertId())

	return ctx.JSON(http.StatusCreated, mhs)
}
