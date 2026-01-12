package user

import (
	"encoding/json"
	"net/http"

	"github.com/kafka6666/ecommerce-crud-gwh/repo"
	"github.com/kafka6666/ecommerce-crud-gwh/utils"
)

type ReqUser struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var reqCreateUser *ReqUser = &ReqUser{}
	if err := json.NewDecoder(r.Body).Decode(reqCreateUser); err != nil {
		utils.SendError(w, http.StatusBadRequest, "Please provide a valid JSON request body")
		return
	}

	user := &repo.User{
		FirstName:   reqCreateUser.FirstName,
		LastName:    reqCreateUser.LastName,
		Email:       reqCreateUser.Email,
		Password:    reqCreateUser.Password,
		IsShopOwner: reqCreateUser.IsShopOwner,
	}

	savedUser, err := h.userRepo.Create(user)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "User not created")
		return
	}

	utils.SendData(w, http.StatusCreated, savedUser)
}
