package product

import (
	"net/http"

	"github.com/kafka6666/ecommerce-crud-gwh/utils"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	productList, err := h.productRepo.ListAll()
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Product list not found")
		return
	}
	utils.SendData(w, http.StatusOK, productList)
}
