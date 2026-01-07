package handlers

import (
	"net/http"

	"github.com/kafka6666/ecommerce-crud-gwh/database"
	"github.com/kafka6666/ecommerce-crud-gwh/utils"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	utils.SendData(w, database.ProductList, http.StatusOK)
}
