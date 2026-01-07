package cmd

import (
	"fmt"
	"net/http"

	"github.com/kafka6666/ecommerce-crud-gwh/middleware"
)

func Serve() {
	// create a default mux as router
	mux := http.NewServeMux()

	// create a middleware manager
	manager := middleware.NewManager()
	manager.Use(middleware.Tester, middleware.Logger, middleware.CorswithPreflight)

	// initialize all routes
	initRoutes(mux, manager)

	// start the server
	fmt.Println("Server running on port:3000")

	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
