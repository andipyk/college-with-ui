package routes

import (
	"college-with-ui/backend/controllers"

	"github.com/labstack/echo"
)

func AdminRoute(group *echo.Group) {
	group.GET("/listdos", controllers.GetAllDosen)     // ambil data Dosen
	group.GET("/listmhs", controllers.GetAllMahasiswa) // ambil data mahasiswa

	group.POST("/inputmhs", controllers.CreateMahasiswa) // tambah mahasiswa
	group.POST("/inputdos", controllers.AddDosen)        // tambah dosen
	group.POST("/inputkls", controllers.AddKelas)        // tambah kelas
}
