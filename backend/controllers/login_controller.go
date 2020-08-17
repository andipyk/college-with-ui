package controllers

import (
	"college-with-ui/backend/database"
	"college-with-ui/backend/models"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func DosenLogin(ctx echo.Context) error {
	db := database.Connect()
	defer db.Close()

	var resDos models.ResponseDosen
	var arrDosen []models.Dosen
	var dosLogin models.Dosen

	if err := ctx.Bind(&dosLogin); err != nil {
		return err
	}
	log.Print(dosLogin.NIK)

	dos := &models.Dosen{}
	err := db.QueryRow("SELECT nik, nama, password FROM dosen WHERE nik = ?", &dosLogin.NIK).Scan(&dos.NIK, &dos.Nama, &dos.Password)
	if err != nil {
		fmt.Println(err)
	}

	// untuk memberi respon ke client
	if dos.NIK == "" {
		resDos.Status = 0
		resDos.Message = "No NIK registered in the system"

		return ctx.JSON(http.StatusBadRequest, resDos)

	} else if dosLogin.NIK != dos.NIK || dosLogin.Password != dos.Password {
		resDos.Status = 0
		resDos.Message = "Please input correct NIK or password"

		return ctx.JSON(http.StatusBadRequest, resDos)

	} else {
		arrDosen = append(arrDosen, models.Dosen{
			NIK:  dos.NIK,
			Nama: dos.Nama,
		})

		resDos.Status = 1
		resDos.Message = "Success Validate Login Data"
		resDos.Data = arrDosen

		return ctx.JSON(http.StatusOK, resDos)
	}
}

func MahasiswaLogin(ctx echo.Context) error {
	db := database.Connect()
	defer db.Close()

	var resMhs models.ResponseMahasiswa
	var arrMhs []models.Mahasiswa
	var mhsLogin models.Mahasiswa

	if err := ctx.Bind(&mhsLogin); err != nil {
		return err
	}
	log.Print(mhsLogin.Nim)

	mhs := &models.Mahasiswa{}
	err := db.QueryRow("SELECT nim, nama, password FROM mahasiswa WHERE nim = ?", &mhsLogin.Nim).Scan(&mhs.Nim, &mhs.Nama, &mhs.Password)
	if err != nil {
		fmt.Println(err)
	}

	// untuk memberi respon ke client
	if mhs.Nim == "" {
		resMhs.Status = 0
		resMhs.Message = "No registered student with that NIM in the system"

		return ctx.JSON(http.StatusBadRequest, resMhs)

	} else if mhsLogin.Nim != mhs.Nim || mhsLogin.Password != mhs.Password {
		resMhs.Status = 0
		resMhs.Message = "Please input correct NIM or password"

		return ctx.JSON(http.StatusBadRequest, resMhs)

	} else {
		arrMhs = append(arrMhs, models.Mahasiswa{
			Nim:  mhs.Nim,
			Nama: mhs.Nama,
		})

		resMhs.Status = 1
		resMhs.Message = "Success Validate Login Data"
		resMhs.Data = arrMhs

		return ctx.JSON(http.StatusOK, resMhs)
	}
}
