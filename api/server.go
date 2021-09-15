package api

import (
	"fmt"
	db "techschool/samplebank/db/sqlc"
	"techschool/samplebank/token"
	"techschool/samplebank/util"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
)

type Server struct {
	config     util.Config // Config struct
	store      db.Store    // type Store interface
	tokenMaker token.Maker // Maker is a interface for managing tokens
	router     *gin.Engine // struct
}

// NewServer creates a http server and setup routing
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetrickey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
		// router: gin.Default(),
	}
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	server.setRouter()
	return server, nil
}

func (server *Server) setRouter() {
	router := gin.Default()

	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)

	// authRoutes: func (group *RouterGroup) Group(relativePath string, handlers ...HandlerFunc) *RouterGroup
	// middleware: HandlerFunc
	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))
	authRoutes.POST("/accounts", server.createAccount)
	authRoutes.GET("/accounts/:id", server.getAccount)
	authRoutes.GET("/accounts", server.listAccounts)

	authRoutes.POST("/transfers", server.createTransfer)

	server.router = router
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
