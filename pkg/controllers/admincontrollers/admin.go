package admincontrollers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/rakhazufar/go-project/pkg/models"
	"github.com/rakhazufar/go-project/pkg/utils"
	"gorm.io/gorm"
)


func AdminLogin(w http.ResponseWriter, r *http.Request) {

}

func AdminRegister(w http.ResponseWriter, r *http.Request) {
	var adminInput models.Admin

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&adminInput); err != nil {
		message := map[string]string {"message": "Failed to decode json"}
		utils.SendJSONResponse(w, http.StatusBadRequest, message)
		return 
	}

	defer r.Body.Close()

	if admin, err := models.GetAdminByUsername(adminInput.Username); err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			message := map[string]string{"message": "Server error"}
			utils.SendJSONResponse(w, http.StatusInternalServerError, message)
			return
		}
	} else if admin != nil {
		message := map[string]string{"message": "Admin Already registered"}
		utils.SendJSONResponse(w, http.StatusConflict, message)
		return 
	}

	if hashedPassword, err := utils.HashAdminPassword(adminInput.Password); err != nil {
		message := map[string]string{"message": "Internal Server Error"}
		utils.SendJSONResponse(w, http.StatusInternalServerError, message)
		return 
	} else {
		adminInput.Password = *hashedPassword
	}
	

	if err := models.CreateAdmin(&adminInput); err != nil {
		message := map[string]string {"message": "Failed to Create User"}
		utils.SendJSONResponse(w, http.StatusInternalServerError, message)
		return 
	}

	message := map[string]string{"message": "success"}
	utils.SendJSONResponse(w, http.StatusOK, message)
}