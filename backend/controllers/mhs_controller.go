package controllers

import (
	"college-with-ui/backend/database"
	"college-with-ui/backend/models"
	"github.com/labstack/echo"
	"log"
	"net/http"
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

func AddMahasiswa(ctx echo.Context) {
	db := database.Connect()
	defer db.Close()

	//err := r.ParseMultipartForm(4096)
	//if err != nil {
	//	panic(err)
	//}

}
