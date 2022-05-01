package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/stankovic004/rezervacija/interfaces"
	"github.com/stankovic004/rezervacija/repo"
)

func RegisterAPI(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var newUser interfaces.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	// fmt.Fprintf(w, "hello, %s email: %s sa lozinkom %s!", newUser.Username, newUser.Email, newUser.Password)
	if err != nil {
		log.Println("can't decode register body: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = repo.Register(newUser)
	if err != nil {
		log.Println("can't register user: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func LoginAPI(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var loginUser interfaces.User
	err := json.NewDecoder(r.Body).Decode(&loginUser)
	// fmt.Fprintf(w, "hello, %s email: %s sa lozinkom %s!", newUser.Username, newUser.Email, newUser.Password)
	if err != nil {
		log.Println("can't decode login body: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = repo.Login(loginUser)
	if err != nil {
		log.Println("can't login user: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}

func ReserveAPI(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var reservation interfaces.Reservation

	err := json.NewDecoder(r.Body).Decode(&reservation)
	if err != nil {
		log.Println("can't reserve this schedule : ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = repo.Reserve(reservation)
	if err != nil {
		log.Println("can't reserve this schedule: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}
