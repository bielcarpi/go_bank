package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"go_bank/db"
	"log"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and sets up routing
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	// Set up a custom validator for currency
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("currency", validCurrency)
		if err != nil {
			log.Fatal("can't set up validator", err)
		}
	}

	// Set up routing
	// We can add middleware to the routes here too
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccounts)
	router.POST("/transfers", server.createTransfer)

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
