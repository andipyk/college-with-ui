package routes

import (
	"college-with-ui/backend/controllers"

	"github.com/labstack/echo"
)

func AdminRoute(group *echo.Group) {
	group.POST("/login", controllers.AdminLogin)

	group.GET("/list_dosen", controllers.GetAllDosen)      // ambil data Dosen
	group.POST("/input/dosen", controllers.AddDosen)       // tambah dosen
	group.PUT("/update/dosen", controllers.EditDosen)      // edit data dosen
	group.DELETE("/delete/dosen", controllers.DeleteDosen) // delete data dosen

	group.GET("/list_mahasiswa", controllers.GetAllMahasiswa)      // ambil data mahasiswa
	group.POST("/input/mahasiswa", controllers.CreateMahasiswa)    // tambah data mahasiswa
	group.PUT("/update/mahasiswa", controllers.EditMahasiswa)      // edit data mahasiswa
	group.DELETE("/delete/mahasiswa", controllers.DeleteMahasiswa) // delete data mahasiswa

	group.GET("/list_kelas", controllers.GetAllKelas)      // ambil data kelas
	group.POST("/input/kelas", controllers.AddKelas)       // tambah kelas
	group.DELETE("/delete/kelas", controllers.DeleteKelas) // delete kelas by kode
}
