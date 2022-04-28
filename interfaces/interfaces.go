package interfaces

import (
	"time"
)

type resComponent struct {
	hour     int
	min      int
	occupied bool
}

type User struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Reservation struct {
	Date     time.Time    `json:"date"`
	Schedule resComponent `json:"schedule"`
}
