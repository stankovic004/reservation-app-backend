package repo

import (
	"errors"
	"fmt"

	"github.com/stankovic004/rezervacija/interfaces"
)

func Register(newUser interfaces.User) error {
	fmt.Println(dbGlobal)
	_, err := dbGlobal.Query(sqlStatements["register"], newUser.Email, newUser.Username, newUser.Password)
	// res, err := dbConn.Exec("insert into users (email, username, password, created_on) VALUES ('mail@gmail.com', 'bravo', 'sifra', NOW());")
	if err != nil {
		fmt.Println(err)
		return err
	}

	return err
}

func Login(loginUser interfaces.User) error {
	fmt.Println(dbGlobal)
	result := 0
	err := dbGlobal.QueryRow(sqlStatements["login"], loginUser.Email, loginUser.Password).Scan(&result)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if result == 0 {
		return errors.New("pogrešni email ili lozinka")
	}

	return err
}

func Reserve(reservation interfaces.Reservation) error {
	fmt.Println(dbGlobal)
	result := 0
	err := dbGlobal.QueryRow(sqlStatements["reserve"], reservation.User, reservation.Date, reservation.Time, reservation.Location).Scan(&result)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if result == 0 {
		return errors.New("greška kod rezervacije")
	}

	return err
}
