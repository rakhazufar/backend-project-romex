package productcontrollers

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gosimple/slug"
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

	u := uuid.New()
	shortUUID := u.String()[:5]
	addressInput.Slug = slug.Make(addressInput.Title + "-" + shortUUID)

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

}

func GetProductsByName(w http.ResponseWriter, r *http.Request) {

}


func EditProductsById(w http.ResponseWriter, r *http.Request) {

}

func DeleteProductsById(w http.ResponseWriter, r *http.Request) {

}