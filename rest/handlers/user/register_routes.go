package user

import (
	"net/http"

	"github.com/kafka6666/ecommerce-crud-gwh/rest/middlewares"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middlewares.Manager) {
	mux.Handle("POST /users", manager.With(http.HandlerFunc(h.CreateUser)))

	mux.Handle("POST /users/login", manager.With(http.HandlerFunc(h.Login)))
}
