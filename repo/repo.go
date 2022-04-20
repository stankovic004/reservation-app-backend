package repo

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
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
