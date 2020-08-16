package controllers

import (
	"college-with-ui/backend/database"
	"college-with-ui/backend/models"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

var dsn models.Dosen
var arrDsn models.DsnArray
var res models.ResponseDosen

func GetAllDosen(ctx echo.Context) error {
	db := database.Connect()
	defer db.Close()

	rows, err := db.Query("SELECT nama, nik FROM dosen")
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&dsn.Nama, &dsn.NIK); err != nil {
			log.Fatal(err.Error())
		} else {
			arrDsn = append(arrDsn, dsn)
		}
	}

	res.Status = 1
	res.Message = "Success"
	res.Data = arrDsn

	// reset kembali splice
	arrDsn = arrDsn[:0]
	return ctx.JSON(http.StatusOK, res)
}

func AddDosen(ctx echo.Context) error {
	db := database.Connect()
	defer db.Close()

	if err := ctx.Bind(dsn); err != nil {
		return err
	}

	query := "INSERT INTO dosen(nama, nik) VALUES (?,?)"
	stmt, err := db.Prepare(query)
	if err != nil {
		fmt.Print(err.Error())
	}

	defer stmt.Close()
	result, err2 := stmt.Exec(dsn.Nama, dsn.NIK)

	// keluar apabila ada error
	if err2 != nil {
		panic(err2)
	}
	fmt.Println(result.LastInsertId())

	return ctx.JSON(http.StatusCreated, res)
}
