package variantcontrollers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rakhazufar/go-project/pkg/config"
	"github.com/rakhazufar/go-project/pkg/models"
	"github.com/rakhazufar/go-project/pkg/utils"
	"gorm.io/gorm"
)

var db = config.GetDB()

func DeleteVariantById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	variantId := vars["id"]
	id, err := strconv.ParseInt(variantId, 10, 64)
	if err != nil {
		log.Printf("Error converting string to int: %v", err)
		message := map[string]string{"message": "Invalid ID format"}
		utils.SendJSONResponse(w, http.StatusInternalServerError, message)
		return
	}

	if err := models.DeleteVariantById(id); err != nil {
		log.Printf("Error deleting image: %v", err)

		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.SendJSONResponse(w, http.StatusNotFound, map[string]string{"message": "Variant not found"})
		} else {
			utils.SendJSONResponse(w, http.StatusInternalServerError, map[string]string{"message": "Internal server error"})
		}
		return
	}

	message := map[string]string{"message": "Success Delete Variant"}
	utils.SendJSONResponse(w, http.StatusOK, message)
}

func GetVariantByProductId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	variantId := vars["id"]
	id, err := strconv.ParseInt(variantId, 10, 64)
	if err != nil {
		log.Printf("Error converting string to int: %v", err)
		message := map[string]string{"message": "Invalid ID format"}
		utils.SendJSONResponse(w, http.StatusInternalServerError, message)
		return
	}

	results, err := models.GetVariantByProductId(id)

	if err != nil {
		message := map[string]string{"message": "Error Get Variant By Products Id"}
		utils.SendJSONResponse(w, http.StatusInternalServerError, message)
		return
	}

	utils.SendJSONResponse(w, http.StatusOK, results)
}

func EditVariant(w http.ResponseWriter, r *http.Request) {
	var variantInput []models.Variant
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&variantInput); err != nil {
		message := map[string]string{"message": "Failed to decode json"}
		utils.SendJSONResponse(w, http.StatusBadRequest, message)
		return
	}
	defer r.Body.Close()

	tx := db.Begin()

	for _, variant := range variantInput {
		result, err := models.GetVariantById(variant.ID)

		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				if err := models.CreateVariant(tx, &variant); err != nil {
					tx.Rollback()
					utils.SendJSONResponse(w, http.StatusInternalServerError, map[string]string{"error": "Could not create variant"})
					return
				}
			} else {
				tx.Rollback()
				utils.SendJSONResponse(w, http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
				return
			}
		}

		if result != nil {
			result.Name = variant.Name
			result.Stock = variant.Stock
			if err := models.UpdateVariant(tx, result); err != nil {
				tx.Rollback()
				utils.SendJSONResponse(w, http.StatusInternalServerError, map[string]string{"error": "Could not update variant"})
				return
			}
		}
	}

	tx.Commit()
	// Kirim respons sukses, misalnya:
	utils.SendJSONResponse(w, http.StatusOK, map[string]string{"message": "Variants updated successfully"})
}
