package product

import (
	"encoding/json"
	"net/http"

	"github.com/kafka6666/ecommerce-crud-gwh/repo"
	"github.com/kafka6666/ecommerce-crud-gwh/utils"
)

type ReqProduct struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"img_url"`
}

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var reqProduct *ReqProduct = &ReqProduct{}
	if err := json.NewDecoder(r.Body).Decode(reqProduct); err != nil {
		utils.SendError(w, http.StatusBadRequest, "Please provide a valid JSON request body")
		return
	}

	newProduct := &repo.Product{
		Title:       reqProduct.Title,
		Description: reqProduct.Description,
		Price:       reqProduct.Price,
		ImgUrl:      reqProduct.ImgUrl,
	}

	savedProduct, err := h.productRepo.Create(newProduct)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Product can not be created.")
		return
	}

	utils.SendData(w, http.StatusCreated, savedProduct)
}
