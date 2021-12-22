package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/timoteoBone/gorilla-sql/app"
	"github.com/timoteoBone/gorilla-sql/db"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := db.CreateDB()
	if err != nil {
		log.Fatal("Db couldnt connect", err.Error())
	}

	app := &app.App{
		Router:   mux.NewRouter().StrictSlash(true),
		Database: db,
	}

	app.SetupRouter()

	s := &http.Server{
		Handler:        app.Router,
		Addr:           ":8000",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()

}
