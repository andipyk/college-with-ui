package routes

import (
	"net/http"

	"github.com/labstack/echo"
)

// return func dari Routes (Echo)
func Index() *echo.Echo {
	e := echo.New()
	e.GET("/", func(context echo.Context) error {
		return context.String(http.StatusOK, "Halaman ini index")
	})

	MahasiswaRoute(e.Group("/mahasiswa"))
	DosenRoute(e.Group("/dosen"))

	return e
}
