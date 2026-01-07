package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/kafka6666/ecommerce-crud-gwh/database"
	"github.com/kafka6666/ecommerce-crud-gwh/utils"
)

type ReqProduct struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"img_url"`
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var reqProduct *ReqProduct = &ReqProduct{}
	if err := json.NewDecoder(r.Body).Decode(reqProduct); err != nil {
		http.Error(w, "Please provide a valid JSON request body", http.StatusBadRequest)
		return
	}

	newProduct := &database.Product{
		ID:          len(database.ProductList) + 1,
		Title:       reqProduct.Title,
		Description: reqProduct.Description,
		Price:       reqProduct.Price,
		ImgUrl:      reqProduct.ImgUrl,
	}

	database.ProductList = append(database.ProductList, newProduct)

	utils.SendData(w, newProduct, http.StatusCreated)
}
