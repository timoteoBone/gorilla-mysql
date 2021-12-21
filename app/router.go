package app

import (
	"database/sql"
	//	"log"
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

}
