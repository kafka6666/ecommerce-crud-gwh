package rest

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/kafka6666/ecommerce-crud-gwh/config"
	"github.com/kafka6666/ecommerce-crud-gwh/rest/middlewares"
)

func Start(cnf *config.Config) {
	// create a default mux as router
	mux := http.NewServeMux()

	// create a middleware manager
	manager := middlewares.NewManager()

	// intiate a global mux/handler wrapped with corswithpreflight middleware
	wrappedMux := manager.WrapMuxWith(
		mux,
		middlewares.Tester,
		middlewares.Logger,
		middlewares.CorswithPreflight)

	// initialize all routes
	initRoutes(mux, manager)

	// start the server
	addr := ":" + strconv.FormatInt(cnf.HttpPort, 10)
	fmt.Println("Server running on port", addr)

	err := http.ListenAndServe(addr, wrappedMux)
	if err != nil {
		fmt.Println("Error starting the server:", err)
		os.Exit(1)
	}
}
