package main

import (
	"github.com/stankovic004/rezervacija/repo"
	"github.com/stankovic004/rezervacija/server"
)

func main() {
	repo.InitConn()
	server.StartServer()
}
