package admincontrollers

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/rakhazufar/go-project/pkg/config"
	"github.com/rakhazufar/go-project/pkg/models"
	"github.com/rakhazufar/go-project/pkg/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)


func AdminLogin(w http.ResponseWriter, r *http.Request) {
	var adminInput models.Admin
	//membaca json dari r.body
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&adminInput); err != nil {
		message := map[string]string{"message":  "Failed to decode json"}
		utils.SendJSONResponse(w, http.StatusBadRequest, message)
		return
	}

	defer r.Body.Close()

	admin, err := models.GetAdminByUsername(adminInput.Username)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			message := map[string]string{"message": "Server error"}
			utils.SendJSONResponse(w, http.StatusInternalServerError, message)
            return
		}
	} else if admin == nil {
		message := map[string]string{"message": "Admin Not Found"}
		utils.SendJSONResponse(w, http.StatusUnauthorized, message)
		return 
	}

	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(adminInput.Password)); err != nil {
		message := map[string]string{"message": "Wrong username/password"}
		utils.SendJSONResponse(w, http.StatusUnauthorized, message)
		return
	}

	expTimeToken := time.Now().Add(time.Hour * 24 * 30)

	claims := &config.AdminJWTClaim{
		AdminId: int(admin.ID),
		RoleID: int(admin.RoleID),
		Username: admin.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: "go-jwt-postgres",
			ExpiresAt: jwt.NewNumericDate(expTimeToken),
		},
	}

	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := tokenAlgo.SignedString(config.JWT_KEY)

	if err != nil {
		message := map[string]string{"message": "Server error"}
		utils.SendJSONResponse(w, http.StatusInternalServerError, message)
        return
	}

	http.SetCookie(w, &http.Cookie{
		Name: "token",
		Path: "/",
		Value: token,
		HttpOnly: true,
	})
	message := map[string]string{"message": "success"}
	utils.SendJSONResponse(w, http.StatusOK, message)
}

func AdminRegister(w http.ResponseWriter, r *http.Request) {
	var adminInput models.Admin

	decoder := json.NewDecoder(r.Body)
	adminInput.RoleID = uint(adminInput.RoleID)
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


func Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name: "token",
		Path: "/",
		Value: "",
		HttpOnly: true,
		MaxAge: -1,
	})
	message := map[string]string{"message": "Logout Success"}
	utils.SendJSONResponse(w, http.StatusOK, message)
}