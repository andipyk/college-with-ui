package routes

import (
	"college-with-ui/backend/controllers"

	"github.com/labstack/echo"
)

func DosenRoute(group *echo.Group) {
	group.GET("/listkelas/:kode", controllers.GetAllKelasByMK) // cek kelas by kode mk
	group.POST("/newdatanilai", controllers.InputNilaiMhs)     // tamnbah nilai
}
