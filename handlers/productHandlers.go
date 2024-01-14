package handlers

import (
	"encoding/json"
	"golang/helpers"
	"golang/models"
	"net/http"
	"strconv"
	"time"
)

var productStore = make(map[string]models.Product)
var id int = 0

// POST - /api/products
func PostProductHandler(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product) //NewDecoder = json çözücü , Decode = body çözümlenir -> çözümlenen veri product adresine yollanır.
	helpers.CheckError(err)

	product.CreatedOn = time.Now()
	id++
	product.ID = id
	keyID := strconv.Itoa(id)
	productStore[keyID] = product

	data, err := json.Marshal(product)
	helpers.CheckError(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(data)
}

// GET - /api/products
func GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	var product []models.Product
	for _, product = range productStore {

	}
}

// GET - /api/products/{id}
func GetProductHandler(w http.ResponseWriter, r *http.Request) {

}

// DELETE - /api/products/{id}
func PutProductHandler(w http.ResponseWriter, r *http.Request) {

}

// DELETE - /api/products/{id}
func DeleteProductHandler(w http.ResponseWriter, r *http.Request) {

}
