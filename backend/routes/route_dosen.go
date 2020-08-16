package routes

import (
	"college-with-ui/backend/controllers"

	"github.com/labstack/echo"
)

func DosenRoute(group *echo.Group) {
	group.GET("", controllers.GetAllDosen)
	//group.POST("/adddsn", controllers.AddMahasiswa) add dosen
	group.POST("/newdatanilai", controllers.InputNilaiMhs)
}
