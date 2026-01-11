package product

import "github.com/kafka6666/ecommerce-crud-gwh/rest/middlewares"

type Handler struct {
	middleware *middlewares.Middleware
}

func NewHandler(m *middlewares.Middleware) *Handler {
	return &Handler{
		middleware: m,
	}
}
