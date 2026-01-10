package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/kafka6666/ecommerce-crud-gwh/database"
	// user "github.com/kafka6666/ecommerce-crud-gwh/database"
	"github.com/kafka6666/ecommerce-crud-gwh/utils"
)

type ReqUser struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var reqCreateUser *ReqUser = &ReqUser{}
	if err := json.NewDecoder(r.Body).Decode(reqCreateUser); err != nil {
		utils.SendError(w, http.StatusBadRequest, "Please provide a valid JSON request body")
		return
	}

	user := &database.User{
		FirstName:   reqCreateUser.FirstName,
		LastName:    reqCreateUser.LastName,
		Email:       reqCreateUser.Email,
		Password:    reqCreateUser.Password,
		IsShopOwner: reqCreateUser.IsShopOwner,
	}

	user.ID = len(user.List()) + 1

	savedUser := user.Store()

	utils.SendData(w, http.StatusCreated, savedUser)
}
