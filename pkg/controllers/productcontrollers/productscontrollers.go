package productcontrollers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rakhazufar/go-project/pkg/models"
	"github.com/rakhazufar/go-project/pkg/utils"
)

func CreateProducts(w http.ResponseWriter, r *http.Request) {
	var addressInput models.Products

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&addressInput); err != nil {
		message := map[string]string{"message":  "Failed to decode json"}
		utils.SendJSONResponse(w, http.StatusBadRequest, message)
		return
	}

	
	addressInput.Slug = utils.Slugify(addressInput.Title)

	defer r.Body.Close()


	if err := models.CreateProduct(&addressInput); err != nil {
		message := map[string]string{"message":  "Internal server error"}
		utils.SendJSONResponse(w, http.StatusInternalServerError, message)
		return
	}

	message := map[string]string{"message":  "Success Add Product"}
		utils.SendJSONResponse(w, http.StatusOK, message)
		return
}

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := models.GetAllProducts(); 
	if err != nil {
		message := map[string]string{"message":  "Internal server error"}
		utils.SendJSONResponse(w, http.StatusInternalServerError, message)
		return
	}

	message := products
	utils.SendJSONResponse(w, http.StatusOK, message)
	return
}

func GetProductsBySlug(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productSlug := vars["slug"]

	product, err := models.GetAllProductsById(productSlug); 
	if err != nil {
		message := map[string]string{"message":  err.Error()}
		utils.SendJSONResponse(w, http.StatusInternalServerError, message)
		return
	}

	message := product
	utils.SendJSONResponse(w, http.StatusOK, message)
	return
}


func EditProductsBySlug(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productSlug := vars["slug"]

	var productInput models.Products
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&productInput); err != nil {
		message := map[string]string{"message":  "Failed to decode json"}
		utils.SendJSONResponse(w, http.StatusBadRequest, message)
		return
	}
	defer r.Body.Close()

	product, err := models.GetAllProductsById(productSlug); 
	if err != nil {
		message := map[string]string{"message":  err.Error()}
		utils.SendJSONResponse(w, http.StatusInternalServerError, message)
		return
	}

	product.Title = productInput.Title
	product.Price = productInput.Price
	product.StatusID =  productInput.StatusID
	product.Description = productInput.Description
	product.Slug = utils.Slugify(productInput.Title)

	if _, err:= models.UpdateProduct(product); err != nil {
		message := map[string]string{"message":  "Failed to updated data"}
		utils.SendJSONResponse(w, http.StatusInternalServerError, message)
		return
	} 

	message := map[string]string{"message":  "Success update Data"}
		utils.SendJSONResponse(w, http.StatusOK, message)
		return
}

func DeleteProductsBySlug(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	slug := vars["slug"]

	if err := models.DeleteProduct(slug); err != nil {
		message := map[string]string{"message":  "Internal server error"}
		utils.SendJSONResponse(w, http.StatusInternalServerError, message)
		return
	}

	message := map[string]string{"message":  "Success Delete product"}
		utils.SendJSONResponse(w, http.StatusOK, message)
		return
}