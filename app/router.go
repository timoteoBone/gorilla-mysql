package app

import (
	"database/sql"

	//"fmt"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	Router   *mux.Router
	Database *sql.DB
}

func (a *App) SetupRouter() {

	a.Router.HandleFunc("/getUser/{id}", a.getUser).Methods("GET")
}

func (a *App) getUser(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		log.Fatal("No ID in the path")
	}

	res, err := a.Database.Query("SELECT id, first_name, last_name FROM USER WHERE id = ?", id) //.Scan(&user.ID, &user.FirstName, &user.LastName)
	if err != nil {
		log.Fatal("db select failed")
	}

	defer res.Close()
	var user User
	for res.Next() {

		err := res.Scan(&user.ID, &user.FirstName, &user.LastName)
		if err != nil {
			panic(err)
		}
		log.Println(user.LastName)
	}

	json.NewEncoder(rw).Encode(user)

}
