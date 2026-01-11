package rest

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/kafka6666/ecommerce-crud-gwh/config"
	productHandler "github.com/kafka6666/ecommerce-crud-gwh/rest/handlers/product"
	userHandler "github.com/kafka6666/ecommerce-crud-gwh/rest/handlers/user"
	"github.com/kafka6666/ecommerce-crud-gwh/rest/middlewares"
)

type Server struct {
	cnf            *config.Config
	userHandler    *userHandler.Handler
	productHandler *productHandler.Handler
	middleware     *middlewares.Middleware
}

func NewServer(
	config *config.Config,
	usrHandler *userHandler.Handler,
	prdctHandler *productHandler.Handler,
	middleware *middlewares.Middleware) *Server {
	return &Server{
		cnf:            config,
		userHandler:    usrHandler,
		productHandler: prdctHandler,
		middleware:     middleware,
	}
}

func (s *Server) Start() {
	// create a default mux as router
	mux := http.NewServeMux()

	// create a middleware manager
	manager := middlewares.NewManager()

	// intiate a global mux/handler wrapped with corswithpreflight middleware
	wrappedMux := manager.WrapMuxWith(
		mux,
		s.middleware.Tester,
		s.middleware.Logger,
		s.middleware.CorswithPreflight)

	// initialize all routes
	s.userHandler.RegisterRoutes(mux, manager)
	s.productHandler.RegisterRoutes(mux, manager)

	// start the server
	addr := ":" + strconv.FormatInt(s.cnf.HttpPort, 10)
	fmt.Println("Server running on port", addr)

	err := http.ListenAndServe(addr, wrappedMux)
	if err != nil {
		fmt.Println("Error starting the server:", err)
		os.Exit(1)
	}
}
