package middleware

import "net/http"

type MiddlewareFn func(http.Handler) http.Handler

type Manager struct {
	globalMiddlewares []MiddlewareFn
}

func NewManager() *Manager {
	return &Manager{
		globalMiddlewares: make([]MiddlewareFn, 0),
	}
}

func (m *Manager) Use(middlewares ...MiddlewareFn) {
	m.globalMiddlewares = append(m.globalMiddlewares, middlewares...)
}

func (m *Manager) With(next http.Handler, middlewares ...MiddlewareFn) http.Handler {
	n := next

	for _, globalMiddleware := range m.globalMiddlewares {
		n = globalMiddleware(n)
	}

	for _, middleware := range middlewares {
		n = middleware(n)
	}

	return n
}
