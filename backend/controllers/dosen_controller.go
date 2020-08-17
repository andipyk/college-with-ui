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
	query := "INSERT INTO dosen(nama, nik, password) VALUES (?,?,?)"
	stmt, err := db.Prepare(query)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer stmt.Close()

	// untuk print di midware
	result, err2 := stmt.Exec(dos.Nama, dos.NIK, dos.Password)
	// keluar apabila ada error
	if err2 != nil {
		panic(err2)
	}
	log.Println(result.LastInsertId())

	// untuk memberi respon ke client
	arrDosen = append(arrDosen, models.Dosen{Nama: dos.Nama, NIK: dos.NIK})

	resDos.Status = 1
	resDos.Message = "Success Add Dosen"
	resDos.Data = arrDosen

	return ctx.JSON(http.StatusCreated, resDos)
}

func EditDosen(ctx echo.Context) error {
	db := database.Connect()
	defer db.Close()

	dos := &models.Dosen{}

	if err := ctx.Bind(dos); err != nil {
		return err
	}
	log.Print(dos)

	result, err := db.Exec("UPDATE dosen SET nama=? WHERE nik = ?",
		dos.Nama,
		dos.NIK,
	)

	if err != nil {
		log.Print(err)
	} else {
		log.Print(result.RowsAffected())
		arrDosen = append(arrDosen, models.Dosen{
			Nama: dos.Nama,
			NIK:  dos.NIK,
		})

	}

	if affected, _ := result.RowsAffected(); affected != 0 {
		resDos.Status = 1
		resDos.Message = "Success Delete Data"
	} else {
		resDos.Status = 2
		resDos.Message = "Data Empty"
	}
	resDos.Data = arrDosen

	return ctx.JSON(http.StatusOK, resDos)
}

func DeleteDosen(ctx echo.Context) error {
	db := database.Connect()
	defer db.Close()

	dos := &models.Dosen{}

	if err := ctx.Bind(dos); err != nil {
		return err
	}
	log.Print(dos.NIK)

	result, err := db.Exec("DELETE FROM dosen where nik = ?",
		dos.Nama,
	)

	if err != nil {
		log.Print(err)
	} else {
		log.Print(result.RowsAffected())
		arrDosen = append(arrDosen, models.Dosen{NIK: dos.NIK})

	}

	if affected, _ := result.RowsAffected(); affected != 0 {
		resDos.Status = 1
		resDos.Message = "Success Delete Data"
	} else {
		resDos.Status = 2
		resDos.Message = "Data Empty"
	}
	resDos.Data = arrDosen

	return ctx.JSON(http.StatusOK, resDos)
}
