package main

import (
	"database/sql"
	"go_bank/api"
	"go_bank/db"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:root@localhost:5432/go_bank?sslmode=disable"
	serverAddress = "localhost:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		panic(err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	// Start the server
	err = server.Start(serverAddress)
	if err != nil {
		panic(err)
	}
}
