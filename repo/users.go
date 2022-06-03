package repo

import (
	"errors"
	"fmt"
	"time"

	"github.com/stankovic004/rezervacija/interfaces"
)

func Register(newUser interfaces.User) error {
	fmt.Println(dbGlobal)
	_, err := dbGlobal.Query(sqlStatements["register"], newUser.Email, newUser.Username, newUser.Password, "user")
	// res, err := dbConn.Exec("insert into users (email, username, password, created_on) VALUES ('mail@gmail.com', 'bravo', 'sifra', NOW());")
	if err != nil {
		fmt.Println(err)
		return err
	}

	return err
}

func Login(loginUser interfaces.User) (interfaces.User, error) {
	fmt.Println(dbGlobal)
	var user  interfaces.User

	err := dbGlobal.QueryRow(sqlStatements["login"], loginUser.Email, loginUser.Password).Scan(&user.Username, &user.Role)
	if err != nil {
		fmt.Println(err)
		return user, err
	}
	if user.Username == "" {
		return user, errors.New("pogrešni email ili lozinka")
	}

	return user, err
}

func Reserve(reservation interfaces.Reservation) error {
	for i := 0; i < len(reservation.Dates); i++ {
		rDate := reservation.Dates[i]
		timeReserved := time.Date(rDate.Year, time.Month(rDate.Month), rDate.Day, rDate.Hour, rDate.Min, 0, 0,time.UTC)
		fmt.Println(timeReserved)
		_, err := dbGlobal.Query(sqlStatements["reserve"], reservation.User, reservation.Location, reservation.Dates)
		if err != nil {
			fmt.Println(err)
			return err
		}

	}
	return nil
}


func AddLocation(location interfaces.Location) error {
	_, err := dbGlobal.Query(sqlStatements["addLocation"], location.Name, location.Lat, location.Lon)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return err
}