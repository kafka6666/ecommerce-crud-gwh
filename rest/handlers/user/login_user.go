package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kafka6666/ecommerce-crud-gwh/config"
	"github.com/kafka6666/ecommerce-crud-gwh/database"
	"github.com/kafka6666/ecommerce-crud-gwh/utils"
)

type ReqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var reqLoginUser *ReqLogin = &ReqLogin{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(reqLoginUser)
	if err != nil {
		fmt.Println(err)
		utils.SendError(w, http.StatusBadRequest, "Please provide a valid json request body")
		return
	}

	foundUser := database.Find(reqLoginUser.Email, reqLoginUser.Password)
	if foundUser == nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid credentials")
		return
	}

	// create a jwt and send it
	cnf := config.GetConfig()

	accessToken, err := utils.CreateJwt(cnf.JwtSecretKey, utils.Payload{
		Sub:         foundUser.ID,
		FirstName:   foundUser.FirstName,
		LastName:    foundUser.LastName,
		Email:       foundUser.Email,
		IsShopOwner: foundUser.IsShopOwner,
	})
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Access token can not be obtained")
		return
	}

	bearerToken := "Bearer " + accessToken

	utils.SendData(w, http.StatusOK, bearerToken)

}
