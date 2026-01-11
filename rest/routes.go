package rest

import (
	"net/http"

	"github.com/kafka6666/ecommerce-crud-gwh/rest/handlers"
	"github.com/kafka6666/ecommerce-crud-gwh/rest/middlewares"
)

func initRoutes(mux *http.ServeMux, manager *middlewares.Manager) {
	// register all routes
	mux.Handle("GET /products", manager.With(http.HandlerFunc(handlers.GetProducts), middlewares.ReceiverLogger))

	mux.Handle("POST /products", manager.With(http.HandlerFunc(handlers.CreateProduct), middlewares.AuthenticateJWT))

	mux.Handle("GET /products/{id}", manager.With(http.HandlerFunc(handlers.GetProductByID)))

	mux.Handle("PUT /products/{id}", manager.With(http.HandlerFunc(handlers.UpdateProduct), middlewares.AuthenticateJWT))

	mux.Handle("DELETE /products/{id}", manager.With(http.HandlerFunc(handlers.DeleteProduct), middlewares.AuthenticateJWT))

	mux.Handle("POST /users", manager.With(http.HandlerFunc(handlers.CreateUser)))

	mux.Handle("POST /users/login", manager.With(http.HandlerFunc(handlers.Login)))
}
