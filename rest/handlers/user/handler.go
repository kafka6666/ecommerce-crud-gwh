package user

import "github.com/kafka6666/ecommerce-crud-gwh/repo"

type Handler struct {
	userRepo repo.UserRepo
}

func NewHandler(repo repo.UserRepo) *Handler {
	return &Handler{
		userRepo: repo,
	}
}
