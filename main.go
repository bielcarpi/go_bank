package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"go_bank/api"
	"go_bank/db"
	"go_bank/util"
	"log"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load configs", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot open connection to DB", err)
	}

	gin.SetMode(gin.DebugMode) //TODO Change to gin.ReleaseMode
	store := db.NewStore(conn)
	server := api.NewServer(store)

	// Start the server
	err = server.Start(config.ServerAddress)
	if err != nil {
		panic(err)
	}
}
