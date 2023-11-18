package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"go_bank/util"
	"log"
	"os"
	"testing"
)

var testQueries *Queries
var testDB *sql.DB

// Entry point for the test
// Opens a connection to the database
func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../")
	if err != nil {
		log.Fatal("cannot load configs", err)
	}

	// Note we can't use := here (we override testDB in global scope)
	testDB, err = sql.Open(config.DBDriver, config.DBSource)

	// Check if the connection is successful
	if err != nil {
		log.Fatal("cannot open connection to DB", err)
	}

	// Assign the connection to the global variable
	testQueries = New(testDB)
	exit := m.Run()
	os.Exit(exit)
}
