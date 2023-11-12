package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
	"testing"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:root@localhost:5432/go_bank?sslmode=disable"
)

var testQueries *Queries
var testDB *sql.DB

// Entry point for the test
// Opens a connection to the database
func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(dbDriver, dbSource)

	// Check if the connection is successful
	if err != nil {
		log.Fatal("Cannot connect to db: ", err)
	}

	// Assign the connection to the global variable
	testQueries = New(testDB)
	exit := m.Run()
	os.Exit(exit)
}
