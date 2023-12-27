package categoriescontrollers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rakhazufar/go-project/pkg/models"
	"github.com/rakhazufar/go-project/pkg/utils"
	"gorm.io/gorm"
)

func DeleteCategoriesById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	categoriesId := vars["id"]
	id, err := strconv.ParseInt(categoriesId, 10, 64)
	if err != nil {
		log.Printf("Error converting string to int: %v", err)
		message := map[string]string{"message": "Invalid ID format"}
		utils.SendJSONResponse(w, http.StatusInternalServerError, message)
		return
	}

	if result, err := models.GetProductCategoryById(id); err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			message := map[string]string{"message": "Server Error"}
			utils.SendJSONResponse(w, http.StatusInternalServerError, message)
			return
		}
	} else if result != nil {
		message := map[string]string{"message": "there is products with this category"}
		utils.SendJSONResponse(w, http.StatusConflict, message)
		return
	}

	if err := models.DeleteCategoryById(id); err != nil {
		log.Printf("Error deleting image: %v", err)

		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.SendJSONResponse(w, http.StatusNotFound, map[string]string{"message": "Category not found"})
		} else {
			utils.SendJSONResponse(w, http.StatusInternalServerError, map[string]string{"message": "Internal server error"})
		}
		return
	}

	message := map[string]string{"message": "Success Delete Category"}
	utils.SendJSONResponse(w, http.StatusOK, message)
}

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	var categoryInput models.Categories
	//membaca json dari r.body
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&categoryInput); err != nil {
		message := map[string]string{"message": "Failed to decode json"}
		utils.SendJSONResponse(w, http.StatusBadRequest, message)
		return
	}
	defer r.Body.Close()

	err := models.CreateCategory(&categoryInput)

	if err != nil {
		message := map[string]string{"message": "Internal Server Error"}
		utils.SendJSONResponse(w, http.StatusInternalServerError, message)
		return
	}

	message := map[string]string{"message": "Success Creating Category"}
	utils.SendJSONResponse(w, http.StatusOK, message)
}

func GetCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := models.GetAllCategories()

	if err != nil {
		message := map[string]string{"message": "Internal Server Error"}
		utils.SendJSONResponse(w, http.StatusInternalServerError, message)
		return
	}

	message := map[string]interface{}{"status": "oke", "data": categories}
	utils.SendJSONResponse(w, http.StatusOK, message)
}
