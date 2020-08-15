package routes

import (
	"college-with-ui/backend/controllers"
	"github.com/labstack/echo"
)

func MahasiswaRoute(group *echo.Group) {
	group.GET("/", controllers.GetAllMahasiswa)
	//group.POST("/", controllers.AddMahasiswa)
}
