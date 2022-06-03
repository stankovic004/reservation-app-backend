package interfaces

type resComponent struct {
	hour     int
	min      int
	occupied bool
	location string
	username string
}

type User struct {
	Email    string `json:"email,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Role string `json:"role,omitempty"`
}

type ReservationDate struct {
	Year int	`json:"year"`
	Month int	`json:"month"`
	Day int		`json:"day"`
	Hour int	`json:"hour"`
	Min int		`json:"min"`
}
type Reservation struct {
	User     string    `json:"user"`
	Location string    `json:"location"`
	Dates     []ReservationDate `json:"dates"`
}

type Location struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Lat  string `json:"lat"`
	Lon  string `json:"lon"`
}
