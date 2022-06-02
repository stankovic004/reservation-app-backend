package repo

import (
	"errors"
	"fmt"

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
