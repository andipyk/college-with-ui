package controllers

import (
	"college-with-ui/backend/database"
	"college-with-ui/backend/models"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

var dos models.Dosen
var arrDosen models.DsnArray
var res models.ResponseDosen

func GetAllDosen(ctx echo.Context) error {
	db := database.Connect()
	defer db.Close()

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

	res.Status = 1
	res.Message = "Success"
	res.Data = arrDosen

	// reset kembali splice
	arrDosen = arrDosen[:0]
	return ctx.JSON(http.StatusOK, res)
}

func AddDosen(ctx echo.Context) error {
	db := database.Connect()
	defer db.Close()

	if err := ctx.Bind(dos); err != nil {
		return err
	}

	query := "INSERT INTO dosen(nama, nik) VALUES (?,?)"
	stmt, err := db.Prepare(query)
	if err != nil {
		fmt.Print(err.Error())
	}

	defer stmt.Close()
	result, err2 := stmt.Exec(dos.Nama, dos.NIK)

	// keluar apabila ada error
	if err2 != nil {
		panic(err2)
	}
	fmt.Println(result.LastInsertId())

	return ctx.JSON(http.StatusCreated, res)
}
