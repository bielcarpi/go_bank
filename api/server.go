package api

import (
	"github.com/gin-gonic/gin"
	"go_bank/db"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and sets up routing
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	// Set up routing
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccounts)

	server.router = router
	return server
}

// Start starts the HTTP server
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

// errorResponse Returns an error response in JSON format to the client
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
