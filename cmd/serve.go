package cmd

import (
	"github.com/kafka6666/ecommerce-crud-gwh/config"
	"github.com/kafka6666/ecommerce-crud-gwh/rest"
)

func Serve() {
	// load configurations from config package
	cnf := config.GetConfig()

	// start the server
	rest.Start(cnf)
}
