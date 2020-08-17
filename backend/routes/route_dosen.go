package routes

import (
	"college-with-ui/backend/controllers"

	"github.com/labstack/echo"
)

func DosenRoute(group *echo.Group) {
	group.GET("/kelas/:kode", controllers.GetAllKelasByMK) // cek kelas by kode mk
	group.POST("/input/nilai", controllers.InputNilaiMhs)  // tamnbah nilai
}
