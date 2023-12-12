package variantcontrollers

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rakhazufar/go-project/pkg/models"
	"github.com/rakhazufar/go-project/pkg/utils"
	"gorm.io/gorm"
)

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
