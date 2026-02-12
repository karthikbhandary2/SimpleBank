package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/karthikbhandary2/simplebank/db/sqlc"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	store *db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccountByID)
	router.GET("/accounts", server.listAccount)
	router.PUT("/account/:id", server.updateAccount)
	router.DELETE("/account/:id", server.deleteAccount)

	//add routes to router
	server.router = router
	return server
}

// Start runs the HTTP server on a specific address
func (server *Server) Start (address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}