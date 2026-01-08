package handlers

import (
	"net/http"
	"strconv"

	"github.com/kafka6666/ecommerce-crud-gwh/database"
	"github.com/kafka6666/ecommerce-crud-gwh/utils"
)

func GetProductByID(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("id")

	id, err := strconv.Atoi(productId)
	if err != nil {
		http.Error(w, "Please provide a valid product id", http.StatusBadRequest)
		return
	}

	product := database.GetByID(id)
	if product == nil {
		utils.SendError(w, http.StatusBadRequest, "Product doesn't exist by the ID provided.")
		return
	}

	utils.SendData(w, http.StatusOK, product)

}
