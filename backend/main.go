package main

import (
	"college-with-ui/backend/routes"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	e := routes.Index()

	e.Logger.Fatal(e.Start(":3000"))
}
