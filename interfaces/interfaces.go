package interfaces

import (
	"time"
)

type resComponent struct {
	hour     int
	min      int
	occupied bool
	location string
	username string
}

type User struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Reservation struct {
	User     string    `json:"user"`
	Time     string    `json:"time"`
	Location string    `json:"location"`
	Occupied bool      `json:"occupied"`
	Date     time.Time `json:"date"`
}

type Location struct {
	Id   int    `json:"id"`
	Name string `json:"email"`
	Lat  string `json:"lat"`
	Lon  string `json:"lon"`
}
