package routes

import (
	"college-with-ui/backend/controllers"

	"github.com/labstack/echo"
)

func MahasiswaRoute(group *echo.Group) {
	group.GET("/mhs/:nim_mhs", controllers.GetMahasiswa) // ambil data mahasiswa by name di tabel nilai
}
