package imagecontrollers

import (
	"context"
	"strconv"

	"log"
	"net/http"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/joho/godotenv"
	"github.com/rakhazufar/go-project/pkg/models"
	"github.com/rakhazufar/go-project/pkg/utils"
)

func UploadImage(w http.ResponseWriter, r *http.Request) {
	
	if err := r.ParseMultipartForm(10 << 20); err != nil {
        utils.SendJSONResponse(w, http.StatusBadRequest, map[string]string{"message": "Error parsing form data"})
        return
    }

	productIDStr := r.FormValue("product_id")
    productID, err := strconv.ParseUint(productIDStr, 10, 32)
    if err != nil {
        http.Error(w, "Invalid product ID", http.StatusBadRequest)
        return
    }

	
	file, handler, err := r.FormFile("image_url")
    if err != nil {
        utils.SendJSONResponse(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
    }
	

	if contentTypes, ok := handler.Header["Content-Type"]; ok && len(contentTypes) > 0 {
		contentType := contentTypes[0]
		if contentType != "image/png" && contentType != "image/jpeg" {
			utils.SendJSONResponse(w, http.StatusBadRequest, map[string]string{"message": "Please upload image png/jpg"})
			return
		}
	} else {
		utils.SendJSONResponse(w, http.StatusBadRequest, map[string]string{"message": "Content-Type header is missing"})
		return
	}
	defer file.Close()
	
    if err = godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    } 

	cloudinaryURL := os.Getenv("CLOUDINARY_URL")

	cldService, _ := cloudinary.NewFromURL(cloudinaryURL)

	ctx := context.Background()

	if result, err := cldService.Upload.Upload(ctx, file, uploader.UploadParams{}); err != nil {
		message := map[string]string{"message":  "Failed to upload image"}
		utils.SendJSONResponse(w, http.StatusInternalServerError, message)
		return
	} else {
		image := models.Image {
			ImageURL:   result.SecureURL,
			ProductID: uint(productID),
		}

		if err := models.ImageUpload(&image); err != nil {
			message := map[string]string{"message":  "failed to save image to database"}
			utils.SendJSONResponse(w, http.StatusInternalServerError, message)
			return
		}

		message := map[string]string{"message":  "success to upload image to database"}
		utils.SendJSONResponse(w, http.StatusOK, message)
		return
	}
	
}
