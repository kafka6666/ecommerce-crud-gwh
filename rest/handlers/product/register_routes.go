package product

import (
	"net/http"

	"github.com/kafka6666/ecommerce-crud-gwh/rest/middlewares"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middlewares.Manager) {
	// register all routes
	mux.Handle("GET /products", manager.With(http.HandlerFunc(h.GetProducts), h.middleware.ReceiverLogger))

	mux.Handle("POST /products", manager.With(http.HandlerFunc(h.CreateProduct), h.middleware.AuthenticateJWT))

	mux.Handle("GET /products/{id}", manager.With(http.HandlerFunc(h.GetProductByID)))

	mux.Handle("PUT /products/{id}", manager.With(http.HandlerFunc(h.UpdateProduct), h.middleware.AuthenticateJWT))

	mux.Handle("DELETE /products/{id}", manager.With(http.HandlerFunc(h.DeleteProduct), h.middleware.AuthenticateJWT))
}
