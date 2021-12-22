package app

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type App struct {
	Router   *mux.Router
	Database *sql.DB
}

func (a *App) SetupRouter() {

	a.Router.HandleFunc("/getUser/{id}", a.getUser).Methods(http.MethodGet)
	a.Router.HandleFunc("/createUser", a.createUser).Methods(http.MethodPost)
	a.Router.HandleFunc("/editUser/{id}", a.editUser).Methods(http.MethodPut)
	a.Router.HandleFunc("/deleteUser/{id}", a.deleteUser).Methods(http.MethodDelete)
}

func (a *App) getUser(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		rw.WriteHeader(http.StatusBadRequest)
	}

	rw.WriteHeader(http.StatusAccepted)

	res, err := a.Database.Query("SELECT id, first_name, last_name FROM USER WHERE id = ?", id)
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
	}

	json.NewEncoder(rw).Encode(user)
}

func (a *App) createUser(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	insert, err := a.Database.Prepare("INSERT INTO USER VALUES(?,?,?)")

	if err != nil {
		log.Fatal(err)
	}

	defer insert.Close()

	user := &User{}

	errr := json.NewDecoder(r.Body).Decode(user)

	if errr != nil {
		log.Fatal(err)
	}

	_, errExc := insert.Exec(&user.ID, &user.FirstName, &user.LastName)

	if errExc != nil {
		log.Fatal(err)
	}
}

func (a *App) editUser(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		rw.WriteHeader(http.StatusBadRequest)
	}

	update, err := a.Database.Prepare("UPDATE USER SET first_name = ?, last_name = ? WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}

	user := &User{}
	json.NewDecoder(r.Body).Decode(user)

	res, errr := update.Exec(&user.FirstName, &user.LastName, id)
	if errr != nil {
		log.Fatal(err)
	}

	rows, errRows := res.RowsAffected()
	if errRows != nil {
		log.Fatal(err)
	}

	log.Println("Rows Affected", strconv.Itoa(int(rows)))
}

func (a *App) deleteUser(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		rw.WriteHeader(http.StatusBadRequest)
	}

	update, err := a.Database.Prepare("DELETE FROM USER WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}

	res, errr := update.Exec(id)
	if errr != nil {
		log.Fatal(err)
	}

	rows, errRows := res.RowsAffected()
	if errRows != nil {
		log.Fatal(err)
	}

	log.Println("Rows Affected", strconv.Itoa(int(rows)))
}
