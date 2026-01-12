package cmd

import (
	"github.com/kafka6666/ecommerce-crud-gwh/config"
	"github.com/kafka6666/ecommerce-crud-gwh/repo"
	"github.com/kafka6666/ecommerce-crud-gwh/rest"
	productHandler "github.com/kafka6666/ecommerce-crud-gwh/rest/handlers/product"
	userHandler "github.com/kafka6666/ecommerce-crud-gwh/rest/handlers/user"
	"github.com/kafka6666/ecommerce-crud-gwh/rest/middlewares"
)

func Serve() {
	// load configurations, handlers and middlewares
	cnf := config.GetConfig()
	middleware := middlewares.NewMiddleware(cnf)
	userRepo := repo.NewUserRepo()
	productRepo := repo.NewProductRepo()
	userHandler := userHandler.NewHandler(userRepo)
	productHandler := productHandler.NewHandler(middleware, productRepo)

	// start the server
	server := rest.NewServer(cnf, userHandler, productHandler, middleware)
	server.Start()
}
