/*
@time: 2021/9/8 13:50
@author: chenZouLu
@file: server
@software: GoLand
@note:
*/

package api

import (
	db "techschool/samplebank/db/sqlc"

	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *db.Store
	router *gin.Engine // struct
}

// NewServer creates a http server and setup routing
func NewServer(store *db.Store) *Server {
	server := &Server{
		store:  store,
		router: gin.Default(),
	}
	// func Default() *Engine {}
	server.router.POST("/accounts", server.CreateAccount())

	return server
}
