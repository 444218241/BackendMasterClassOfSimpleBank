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

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
)

type Server struct {
	store  db.Store    // type Store interface
	router *gin.Engine // struct
}

// NewServer creates a http server and setup routing
func NewServer(store db.Store) *Server {
	server := &Server{
		store:  store,
		router: gin.Default(),
	}
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	server.router.POST("/user", server.createUser)
	server.router.POST("/transfer", server.createTransfer)

	server.router.POST("/account", server.createAccount)
	server.router.GET("/account/:id", server.getAccount)
	server.router.GET("/accounts", server.listAccounts)

	return server
}

// gin.H is short for  type H map[string]interface{}
// type error interface {
//	Error() string
// }
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (server *Server) Start(addr string) error {
	return server.router.Run(addr)
}
