package addresscontrollers

import (
	"encoding/json"
	"errors"

	"net/http"

	"github.com/rakhazufar/go-project/pkg/models"
	"github.com/rakhazufar/go-project/pkg/utils"
	"gorm.io/gorm"
)

func CreateAddress(w http.ResponseWriter, r *http.Request) {
	var addressInput models.Address
	//membaca json dari r.body
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&addressInput); err != nil {
		message := map[string]string{"message":  "Failed to decode json"}
		utils.SendJSONResponse(w, http.StatusBadRequest, message)
		return
	}
	defer r.Body.Close()

	username, ok := r.Context().Value("username").(string)
	
	if !ok {
        message := map[string]string{"message":  "Unauthorized"}
		utils.SendJSONResponse(w, http.StatusUnauthorized, message)
		return
    }

	result, err := models.GetUserByUsername(username);
	
	if err != nil {
		message := map[string]string{"message":  "Token invalid"}
		utils.SendJSONResponse(w, http.StatusUnauthorized, message)
		return
	}

	address, err := models.GetAddressByUserId(int(result.ID))

	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			message := map[string]string{"message": "Server Error"}
			utils.SendJSONResponse(w, http.StatusInternalServerError, message)
            return
		}
	} else if address != nil {
		message := map[string]string{"message": "You already have address"}
		utils.SendJSONResponse(w, http.StatusBadRequest, message)
    	return
	}

	addressInput.User = *result
	addressInput.User.Username = result.Username

	if err := models.CreateAddress(&addressInput); err != nil {
		message := map[string]string{"message":  "Internal server error"}
		utils.SendJSONResponse(w, http.StatusInternalServerError, message)
		return
	}

	message := map[string]string{"message":  "Success Creating Address"}
		utils.SendJSONResponse(w, http.StatusOK, message)
		return
}

func GetAddress(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value("userId").(int)
	
	if !ok {
        message := map[string]string{"message":  "Unauthorized"}
		utils.SendJSONResponse(w, http.StatusUnauthorized, message)
		return
    }

	if userAddress, err := models.GetAddressByUserId(userId); err != nil {
		message := map[string]string{"message":  "Cannot find Address"}
		utils.SendJSONResponse(w, http.StatusNotFound, message)
		return
	} else {
		message := userAddress
		utils.SendJSONResponse(w, http.StatusOK, message)
		return
	}
}

func UpdateAddress(w http.ResponseWriter, r *http.Request) {
	var addressInput models.Address
	//membaca json dari r.body
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&addressInput); err != nil {
		message := map[string]string{"message":  "Failed to decode json"}
		utils.SendJSONResponse(w, http.StatusBadRequest, message)
		return
	}
	defer r.Body.Close()

	userId, ok := r.Context().Value("userId").(int)
	
	if !ok {
        message := map[string]string{"message":  "Unauthorized"}
		utils.SendJSONResponse(w, http.StatusUnauthorized, message)
		return
    }

	userAddress, err := models.GetAddressByUserId(userId);

	if err != nil {
		message := map[string]string{"message":  "Cannot find Address"}
		utils.SendJSONResponse(w, http.StatusNotFound, message)
		return
	}

	userAddress.City = addressInput.City
	userAddress.State = addressInput.State
	userAddress.Country = addressInput.Country
	userAddress.FullAddress = addressInput.FullAddress
	userAddress.PostalCode = addressInput.PostalCode
	
	if _, err:= models.UpdateAddress(userAddress); err != nil {
		message := map[string]string{"message":  "Failed to updated data"}
		utils.SendJSONResponse(w, http.StatusInternalServerError, message)
		return
	} 

	message := map[string]string{"message":  "Success update Data"}
		utils.SendJSONResponse(w, http.StatusOK, message)
		return
}