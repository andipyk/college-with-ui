package controllers

import (
	"college-with-ui/backend/database"
	"college-with-ui/backend/models"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

var arrDosen models.DosArray
var resDos models.ResponseDosen

func GetAllDosen(ctx echo.Context) error {
	db := database.Connect()
	defer db.Close()

	var dos models.Dosen

	rows, err := db.Query("SELECT nama, nik FROM dosen")
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&dos.Nama, &dos.NIK); err != nil {
			log.Fatal(err.Error())
		} else {
			arrDosen = append(arrDosen, dos)
		}
	}

	resDos.Status = 1
	resDos.Message = "Success"
	resDos.Data = arrDosen

	// reset kembali splice
	arrDosen = arrDosen[:0]
	return ctx.JSON(http.StatusOK, resDos)
}

func AddDosen(ctx echo.Context) error {
	db := database.Connect()
	defer db.Close()

	// bind struct dengan message body
	dos := &models.Dosen{}
	if err := ctx.Bind(dos); err != nil {
		return err
	}

	// buat query untuk DB
	query := "INSERT INTO dosen(nama, nik) VALUES (?,?)"
	stmt, err := db.Prepare(query)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer stmt.Close()

	// untuk print di midware
	result, err2 := stmt.Exec(dos.Nama, dos.NIK)
	// keluar apabila ada error
	if err2 != nil {
		panic(err2)
	}
	fmt.Println(result.LastInsertId())

	// untuk memberi respon ke client
	arrDosen = append(arrDosen, models.Dosen{Nama: dos.Nama, NIK: dos.NIK})

	resDos.Status = 1
	resDos.Message = "Success Add Dosen"
	resDos.Data = arrDosen

	return ctx.JSON(http.StatusCreated, resDos)
}
