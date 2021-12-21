package app

import (
	"database/sql"
	"log"

	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	Router   *mux.Router
	Database sql.DB
}

func (a *App) SetupRouter() {

	a.Router.Methods(http.MethodGet).Path("/getUser/{id}").HandlerFunc(a.getUser)
}

func (a *App) getUser(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		log.Fatal("No ID in the path")
	}

	//TO DO
	//data,err:=a.Database.Query("SELECT first_name FROM USER WHERE id = ")
	//still dont know how to show the data to the request
}
