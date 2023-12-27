package statuscontrollers

import (
	"net/http"

	"github.com/rakhazufar/go-project/pkg/models"
	"github.com/rakhazufar/go-project/pkg/utils"
)

func GetStatus(w http.ResponseWriter, r *http.Request) {
	status, err := models.GetAllStatus()

	if err != nil {
		message := map[string]string{"message": "Internal Server Error"}
		utils.SendJSONResponse(w, http.StatusInternalServerError, message)
		return
	}

	message := map[string]interface{}{"status": "oke", "data": status}
	utils.SendJSONResponse(w, http.StatusOK, message)
}
