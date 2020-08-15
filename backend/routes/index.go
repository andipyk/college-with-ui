package routes

import (
	"github.com/labstack/echo"
	"net/http"
)

// return func dari Routes (Echo)
func Index() *echo.Echo {
	e := echo.New()
	e.GET("/", func(context echo.Context) error {
		return context.String(http.StatusOK, "Halaman ini index")
	})

	MahasiswaRoute(e.Group("/mahasiswa"))

	return e
}
