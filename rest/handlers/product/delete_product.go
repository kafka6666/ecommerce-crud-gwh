package product

import (
	"net/http"
	"strconv"

	"github.com/kafka6666/ecommerce-crud-gwh/database"
	"github.com/kafka6666/ecommerce-crud-gwh/utils"
)

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("id")

	id, err := strconv.Atoi(productId)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Please provide a valid product id")
		return
	}

	err = database.Delete(id)
	if err != nil {
		utils.SendError(w, http.StatusNotFound, "Product not found")
		return
	}

	utils.SendData(w, http.StatusCreated, "Product deleted successfully")
}
