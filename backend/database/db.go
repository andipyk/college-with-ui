package database

import (
	"database/sql"
	"log"
)

func Connect() *sql.DB {
	const (
		user       = "1pn4yZkAf8"
		pass       = "MKjwE7Xitv"
		server     = "remotemysql.com"
		port       = "3306"
		dbname     = "1pn4yZkAf8"
		datasource = user + ":" + pass + "@tcp(" + server + ":" + port + ")/" + dbname
	)

	db, err := sql.Open("mysql", datasource)

	if err != nil {
		log.Fatal(err)
	}

	return db
}
