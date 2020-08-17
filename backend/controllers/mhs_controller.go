package controllers

import (
	"college-with-ui/backend/database"
	"college-with-ui/backend/models"
	"fmt"
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
	response.Message = "Success Add Mahasiswa"
	response.Data = arrMhs

	log.Print("Insert data to database")

	return ctx.JSON(http.StatusCreated, response)
}

func EditMahasiswa(ctx echo.Context) error {
	db := database.Connect()
	defer db.Close()

	var arrMhs []models.Mahasiswa
	mhs := &models.Mahasiswa{}

	if err := ctx.Bind(mhs); err != nil {
		return err
	}
	log.Print(mhs)

	result, err := db.Exec("UPDATE mahasiswa SET nama=? WHERE nim = ?",
		mhs.Nama,
		mhs.Nim,
	)

	if err != nil {
		log.Print(err)
	} else {
		log.Print(result.RowsAffected())
		arrMhs = append(arrMhs, models.Mahasiswa{
			Nama: mhs.Nama,
			Nim:  mhs.Nim,
		})

	}

	if affected, _ := result.RowsAffected(); affected != 0 {
		response.Status = 1
		response.Message = "Success Delete Data"
	} else {
		response.Status = 2
		response.Message = "Data Empty"
	}
	response.Data = arrMhs

	return ctx.JSON(http.StatusOK, response)
}

func DeleteMahasiswa(ctx echo.Context) error {
	db := database.Connect()
	defer db.Close()

	var arrMhs []models.Mahasiswa
	mhs := &models.Mahasiswa{}

	if err := ctx.Bind(mhs); err != nil {
		return err
	}
	log.Print(mhs.Nim)

	result, err := db.Exec("DELETE FROM mahasiswa where nim = ?",
		mhs.Nim,
	)

	if err != nil {
		log.Print(err)
	} else {
		log.Print(result.RowsAffected())
		arrMhs = append(arrMhs, models.Mahasiswa{Nim: mhs.Nim})

	}

	if affected, _ := result.RowsAffected(); affected != 0 {
		response.Status = 1
		response.Message = "Success Delete Data"
	} else {
		response.Status = 2
		response.Message = "Data Empty"
	}
	response.Data = arrMhs

	return ctx.JSON(http.StatusOK, response)
}

func GetMahasiswa(ctx echo.Context) error {
	var arrReport models.ReportArray
	var nilai models.Nilai
	var resRep models.ResponseReport

	db := database.Connect()
	defer db.Close()

	requestedNim := ctx.Param("nim_mhs")
	err := db.QueryRow("SELECT nim_mhs, nama_mhs, nik_dosen, kode_mk, nama_mk, absen, nilai, total_nilai FROM nilai WHERE nim_mhs = ?", requestedNim).Scan(&nilai.NIM, &nilai.NamaMhs, &nilai.NIKDosen, &nilai.KodeMK, &nilai.NamaMK, &nilai.Absen, &nilai.Nilai, &nilai.TotalNilai)
	if err != nil {
		fmt.Println(err)
	}

	// untuk memberi respon ke client
	arrReport = append(arrReport, models.Nilai{
		NIM:        nilai.NIM,
		NamaMhs:    nilai.NamaMhs,
		NIKDosen:   nilai.NIKDosen,
		KodeMK:     nilai.KodeMK,
		NamaMK:     nilai.NamaMK,
		Absen:      nilai.Absen,
		Nilai:      nilai.Nilai,
		TotalNilai: nilai.TotalNilai,
	})

	resRep.Status = 1
	resRep.Message = "Success Fetch Data"
	resRep.Data = arrReport

	return ctx.JSON(http.StatusOK, resRep)
}
