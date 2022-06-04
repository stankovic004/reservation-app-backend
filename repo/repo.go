package repo

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/stankovic004/rezervacija/interfaces"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "rezervacija"
)

var dbGlobal *sql.DB

func InitConn() error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	dbConn, err := sql.Open("postgres", psqlInfo)
	dbGlobal = dbConn

	if err != nil {
		fmt.Printf("something went wrong with db: %v\n", err)
		return nil
	}

	err = dbConn.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
	return nil
}

func GetLocations() ([]interfaces.Location, error) {
	var locations []interfaces.Location
	rows, err := dbGlobal.Query(sqlStatements["get_locations"])
	if err != nil {
		fmt.Println(err)
		return locations, err
	}

	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		var lon string
		var lat string
		err = rows.Scan(&id, &name, &lon, &lat)
		if err != nil {
			// handle this error
			fmt.Println(err)
		}
		l := interfaces.Location{id, name, lat, lon}
		locations = append(locations, l)
	}
	return locations, err
}

func GetReservations() ([]interfaces.GetReservationType, error) {
	var reservations []interfaces.GetReservationType
	rows, err := dbGlobal.Query(sqlStatements["get_reservations"])
	if err != nil {
		fmt.Println(err)
		return reservations, err
	}

	defer rows.Close()
	for rows.Next() {
		var id int
		var user string
		var location string
		var dates string
		err = rows.Scan(&id, &user, &location, &dates)
		if err != nil {
			// handle this error
			fmt.Println(err)
		}
		l := interfaces.GetReservationType{id, user, location, dates}
		reservations = append(reservations, l)
	}
	return reservations, err
}
