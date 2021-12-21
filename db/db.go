package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func CreateDB() (*sql.DB, error) {

	db, err := sql.Open("mysql", "root:PewDiePie8!!@tcp(127.0.0.1:3036)/test")

	if err != nil {
		return nil, err
	}

	return db, nil
}
