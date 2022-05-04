package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"github.com/stankovic004/rezervacija/repo"
)

func StartServer() {
	router := httprouter.New()
	router.GET("/hello/:name", Hello)
	router.POST("/register", RegisterAPI)
	router.POST("/login", LoginAPI)
	router.GET("/locations", GetLocationsAPI)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "http://foo.com:8080", "."},
		AllowCredentials: true,
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	})
	handler := c.Handler(router)
	log.Fatal(http.ListenAndServe(":8080", handler))
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func GetLocationsAPI(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	locations, err := repo.GetLocations()
	if err != nil {
		log.Println("error: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println(locations)
	err = json.NewEncoder(w).Encode(locations)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return

	}
	w.WriteHeader(http.StatusOK)
}
