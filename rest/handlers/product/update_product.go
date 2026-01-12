package product

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/kafka6666/ecommerce-crud-gwh/repo"
	"github.com/kafka6666/ecommerce-crud-gwh/utils"
)

type ReqUpdatedProduct struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"img_url"`
}

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("id")

	id, err := strconv.Atoi(productId)
	if err != nil {
		http.Error(w, "Please provide a valid product id", http.StatusBadRequest)
		return
	}

	var reqProduct *ReqUpdatedProduct = &ReqUpdatedProduct{}
	if err := json.NewDecoder(r.Body).Decode(reqProduct); err != nil {
		http.Error(w, "Please provide a valid JSON request body", http.StatusBadRequest)
		return
	}

	updatedProduct := &repo.Product{
		ID:          id,
		Title:       reqProduct.Title,
		Description: reqProduct.Description,
		Price:       reqProduct.Price,
		ImgUrl:      reqProduct.ImgUrl,
	}

	savedProduct, err := h.productRepo.Update(updatedProduct)
	if err != nil {
		utils.SendError(w, http.StatusNotFound, "Product can not be saved")
		return
	}

	log.Println(savedProduct)

	utils.SendData(w, http.StatusOK, "Succesfully updated the product")
}
