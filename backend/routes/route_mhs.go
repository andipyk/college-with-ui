package routes

import (
	"college-with-ui/backend/controllers"

	"github.com/labstack/echo"
)

func MahasiswaRoute(group *echo.Group) {
	group.POST("/login", controllers.MahasiswaLogin)
	group.GET("/report/:nim_mhs", controllers.GetMahasiswa) // ambil data mahasiswa by name di tabel nilai
}
