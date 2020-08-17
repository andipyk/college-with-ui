package routes

import (
	"college-with-ui/backend/controllers"

	"github.com/labstack/echo"
)

func AdminRoute(group *echo.Group) {
	group.GET("/list", controllers.GetAllDosen) // ambil data Dosen
	group.POST("/input", controllers.AddDosen)  // tambah dosen

	group.GET("/list", controllers.GetAllMahasiswa)      // ambil data mahasiswa
	group.POST("/input", controllers.CreateMahasiswa)    // tambah data mahasiswa
	group.PUT("/update", controllers.EditMahasiswa)      // edit data mahasiswa
	group.DELETE("/delete", controllers.DeleteMahasiswa) // delete data mahasiswa

	group.POST("/inputkls", controllers.AddKelas) // tambah kelas
}
