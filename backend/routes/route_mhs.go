package routes

import (
	"college-with-ui/backend/controllers"

	"github.com/labstack/echo"
)

func MahasiswaRoute(group *echo.Group) {
	group.GET("", controllers.GetAllMahasiswa)
	group.POST("", controllers.CreateMahasiswa)
	group.PUT("", controllers.EditMahasiswa)
	group.DELETE("", controllers.DeleteMahasiswa)
	group.GET("/report/:nim_mhs", controllers.GetMahasiswa) // ambil data mahasiswa by name di tabel nilai
}
