package api

import (
	db "github.com/afiifatuts/simple_bank/db/sqlc"
	"github.com/gin-gonic/gin"
)

// Server serves HTTP request for our banking service
type Server struct {
	store  db.Store
	router *gin.Engine
}

// Newserver creates a new HTTO serber and setup routing
func NewServer(strore db.Store) *Server {
	server := &Server{store: strore}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)

	server.router = router
	return server
}

// Start runs the HTTP server on a specific router
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
