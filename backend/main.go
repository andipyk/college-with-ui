package main

import (
	"github.com/labstack/echo"
	"net/http"
)

type User struct {
	Name  string `json:"name" xml:"name" form:"name" query:"name"`
	Email string `json:"email" xml:"email" form:"email" query:"email"`
}

func main() {
	e := echo.New()
	e.GET("/", func(context echo.Context) error {
		return context.String(http.StatusOK, "halo dunia")
	})

	e.POST("/users", saveUser)

	e.Logger.Fatal(e.Start(":3000"))
}



// e.POST("/save", save)
func saveUser(context echo.Context) error {
	u := new(User)
	if err := context.Bind(u); err != nil {
		return err
	}
	return context.JSON(http.StatusCreated, u)
}
