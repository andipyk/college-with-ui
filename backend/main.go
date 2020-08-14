package main

import (
	"college-with-ui/backend/models"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"log"
	"net/http"
)

func main() {
	var mhs models.Mahasiswa
	var arrMhs []models.Mahasiswa
	var response models.ResponseMahasiswa

	db := connect()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM mahasiswa")
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

	e := echo.New()
	e.GET("/", func(context echo.Context) error {
		//return context.String(http.StatusOK, "halo dunia")
		return context.JSON(http.StatusOK, response)
	})

	e.Logger.Fatal(e.Start(":3000"))
}

func connect() *sql.DB {
	const (
		user       = "1pn4yZkAf8"
		pass       = "MKjwE7Xitv"
		server     = "remotemysql.com"
		port       = "3306"
		dbname     = "1pn4yZkAf8"
		datasource = user + ":" + pass + "@tcp(" + server + ":" + port + ")/" + dbname
	)

	db, err := sql.Open("mysql", datasource)

	if err != nil {
		log.Fatal(err)
	}

	return db
}
