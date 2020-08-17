package routes

import (
	"college-with-ui/backend/controllers"

	"github.com/labstack/echo"
)

func DosenRoute(group *echo.Group) {
	group.POST("/login", controllers.DosenLogin)
	group.GET("/kelas/:kode", controllers.GetAllKelasByMK) // cek kelas by kode mk, untuk lihat report
	group.POST("/input/nilai", controllers.InputNilaiMhs)  // tamnbah nilai
}
