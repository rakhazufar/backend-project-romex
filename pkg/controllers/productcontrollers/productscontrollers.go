package productcontrollers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rakhazufar/go-project/pkg/config"
	"github.com/rakhazufar/go-project/pkg/models"
	"github.com/rakhazufar/go-project/pkg/utils"
)

var db = config.GetDB()

func CreateProducts(w http.ResponseWriter, r *http.Request) {
	var input models.ProductWithVariantsInput

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&input); err != nil {
		message := map[string]string{"message": "Failed to decode json"}
		utils.SendJSONResponse(w, http.StatusBadRequest, message)
		return
	}

	input.Product.Slug = utils.Slugify(input.Product.Title)

	defer r.Body.Close()

	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	fmt.Print(input.Product)
	product, err := models.CreateProduct(tx, &input.Product)

	if err != nil {
		tx.Rollback()
		utils.SendJSONResponse(w, http.StatusInternalServerError, map[string]string{"message": "Internal server error"})
		return
	}

	for _, variantInput := range input.Variants {
		variantInput.ProductsID = int(product.ID)
		if err := models.CreateVariant(tx, &variantInput); err != nil {
			tx.Rollback()
			utils.SendJSONResponse(w, http.StatusInternalServerError, map[string]string{"error": "Could not create variant"})
			return
		}
	}

	if err := tx.Commit().Error; err != nil {
		utils.SendJSONResponse(w, http.StatusInternalServerError, map[string]string{"error": "Transaction commit failed"})
		return
	}

	message := map[string]string{"message": "Product and variants created successfully"}
	utils.SendJSONResponse(w, http.StatusOK, message)
}

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := models.GetAllProducts()
	if err != nil {
		message := map[string]string{"message": "Internal server error"}
		utils.SendJSONResponse(w, http.StatusInternalServerError, message)
		return
	}

	message := products
	utils.SendJSONResponse(w, http.StatusOK, message)
}

func GetProductsBySlug(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productSlug := vars["slug"]

	product, err := models.GetProductBySlug(productSlug)
	if err != nil {
		message := map[string]string{"message": err.Error()}
		utils.SendJSONResponse(w, http.StatusInternalServerError, message)
		return
	}

	message := product
	utils.SendJSONResponse(w, http.StatusOK, message)
}

func EditProductsBySlug(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productSlug := vars["slug"]

	var productInput models.Products
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&productInput); err != nil {
		message := map[string]string{"message": "Failed to decode json"}
		utils.SendJSONResponse(w, http.StatusBadRequest, message)
		return
	}
	defer r.Body.Close()

	product, err := models.GetProductBySlug(productSlug)
	if err != nil {
		message := map[string]string{"message": err.Error()}
		utils.SendJSONResponse(w, http.StatusInternalServerError, message)
		return
	}

	product.Title = productInput.Title
	product.Price = productInput.Price
	product.StatusID = productInput.StatusID
	product.Description = productInput.Description
	product.Slug = utils.Slugify(productInput.Title)

	if _, err := models.UpdateProduct(product); err != nil {
		message := map[string]string{"message": "Failed to updated data"}
		utils.SendJSONResponse(w, http.StatusInternalServerError, message)
		return
	}

	message := map[string]string{"message": "Success update Data"}
	utils.SendJSONResponse(w, http.StatusOK, message)
}

func DeleteProductsBySlug(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	slug := vars["slug"]

	if err := models.DeleteProduct(slug); err != nil {
		message := map[string]string{"message": "Internal server error"}
		utils.SendJSONResponse(w, http.StatusInternalServerError, message)
		return
	}

	message := map[string]string{"message": "Success Delete product"}
	utils.SendJSONResponse(w, http.StatusOK, message)
}
