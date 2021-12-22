package db

import (
	"database/sql"
	//"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func CreateDB() (*sql.DB, error) {

	db, err := sql.Open("mysql", "root:PewDiePie8!!@tcp(127.0.0.1:3306)/test?parseTime=true")

	if err != nil {
		panic(err.Error())
	}

	return db, nil
}
