package routes

import (
	"college-with-ui/backend/controllers"

	"github.com/labstack/echo"
)

func MahasiswaRoute(group *echo.Group) {
	group.GET("", controllers.GetAllMahasiswa)
	group.POST("/addmhs", controllers.AddMahasiswa)
}

func DosenRoute(group *echo.Group) {
	group.GET("", controllers.GetAllDosen)
	group.POST("/adddsn", controllers.AddMahasiswa)
	group.POST("/newdatanilai", controllers.InputNilaiMhs)
}
