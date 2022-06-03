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
	user, err := repo.Login(loginUser)
	if err != nil {
		log.Println("can't login user: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		log.Println("can't encode user: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func AddReservationsAPI(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var reservation interfaces.Reservation
	err := json.NewDecoder(r.Body).Decode(&reservation)
	if err != nil {
		log.Println("can't decode this schedule : ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Println(reservation)
	err = repo.Reserve(reservation)
	if err != nil {
		log.Println("can't reserve this schedule: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}

func AddLocationAPI(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	
	var location interfaces.Location
	
	err := json.NewDecoder(r.Body).Decode(&location)
	if err != nil {
		log.Println("can't add location ", err)
		log.Println(location.Name)
		log.Println(location.Lon)
		log.Println(location.Lat)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = repo.AddLocation(location)
	if err != nil {
		log.Println(location.Name)
		log.Println(location.Lon)
		log.Println(location.Lat)
		log.Println("can't add location : ", err)
		
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
