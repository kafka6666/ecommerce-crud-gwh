package rest

import (
	"net/http"

	"github.com/kafka6666/ecommerce-crud-gwh/rest/handlers"
	"github.com/kafka6666/ecommerce-crud-gwh/rest/middlewares"
)

func initRoutes(mux *http.ServeMux, manager *middlewares.Manager) {
	// register all routes
	mux.Handle("GET /products", manager.With(http.HandlerFunc(handlers.GetProducts), middlewares.ReceiverLogger))

	mux.Handle("POST /products", manager.With(http.HandlerFunc(handlers.CreateProduct)))

	mux.Handle("GET /products/{id}", manager.With(http.HandlerFunc(handlers.GetProductByID)))

	mux.Handle("PUT /products/{id}", manager.With(http.HandlerFunc(handlers.UpdateProduct)))

	mux.Handle("DELETE /products/{id}", manager.With(http.HandlerFunc(handlers.DeleteProduct)))
}
