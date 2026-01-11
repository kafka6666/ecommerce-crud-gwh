package cmd

import (
	"github.com/kafka6666/ecommerce-crud-gwh/config"
	"github.com/kafka6666/ecommerce-crud-gwh/rest"
	productHandler "github.com/kafka6666/ecommerce-crud-gwh/rest/handlers/product"
	userHandler "github.com/kafka6666/ecommerce-crud-gwh/rest/handlers/user"
	"github.com/kafka6666/ecommerce-crud-gwh/rest/middlewares"
)

func Serve() {
	// load configurations, handlers and middlewares
	cnf := config.GetConfig()
	middleware := middlewares.NewMiddleware(cnf)
	userHandler := userHandler.NewHandler()
	productHandler := productHandler.NewHandler(middleware)

	// start the server
	server := rest.NewServer(cnf, userHandler, productHandler, middleware)
	server.Start()
}
