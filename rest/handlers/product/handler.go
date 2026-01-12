package product

import (
	"github.com/kafka6666/ecommerce-crud-gwh/repo"
	"github.com/kafka6666/ecommerce-crud-gwh/rest/middlewares"
)

type Handler struct {
	middleware  *middlewares.Middleware
	productRepo repo.ProductRepo
}

func NewHandler(m *middlewares.Middleware, repo repo.ProductRepo) *Handler {
	return &Handler{
		middleware:  m,
		productRepo: repo,
	}
}
