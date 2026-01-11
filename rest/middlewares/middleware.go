package middlewares

import "github.com/kafka6666/ecommerce-crud-gwh/config"

type Middleware struct {
	cnf *config.Config
}

func NewMiddleware(config *config.Config) *Middleware {
	return &Middleware{
		cnf: config,
	}
}
