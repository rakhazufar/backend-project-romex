package imagecontrollers

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"net/http"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gorilla/mux"
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

	cloudinaryURL := os.Getenv("CLOUDINARY_URL")

	cldService, _ := cloudinary.NewFromURL(cloudinaryURL)

	files := r.MultipartForm.File["image_url"]

	for x, fileHeader := range files {
		file, err := fileHeader.Open()

		fmt.Printf(fileHeader.Filename)
		if err != nil {
			utils.SendJSONResponse(w, http.StatusInternalServerError, map[string]string{"message": "Error opening file"})
			continue
		}

		if contentTypes, ok := fileHeader.Header["Content-Type"]; ok && len(contentTypes) > 0 {
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

		ctx := context.Background()

		if result, err := cldService.Upload.Upload(ctx, file, uploader.UploadParams{}); err != nil {
			message := map[string]string{"message": "Failed to upload image"}
			utils.SendJSONResponse(w, http.StatusInternalServerError, message)
			return
		} else {
			fmt.Printf(`Terpanggil %v`, x)
			image := models.Image{
				ImageURL:  result.SecureURL,
				ProductID: uint(productID),
			}

			if err := models.ImageUpload(&image); err != nil {
				message := map[string]string{"message": "failed to save image to database"}
				utils.SendJSONResponse(w, http.StatusInternalServerError, message)
				return
			}
		}
	}

	message := map[string]string{"message": "success to upload image to database"}
	utils.SendJSONResponse(w, http.StatusOK, message)
}

func GetImagesBySlug(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productSlug := vars["slug"]

	products, err := models.GetProductBySlug(productSlug)

	if err != nil {
		message := map[string]string{"message": err.Error()}
		utils.SendJSONResponse(w, http.StatusConflict, message)
		return
	}

	if images, err := models.GetImages(int(products.ID)); err != nil {
		message := map[string]string{"message": "failed to get images"}
		utils.SendJSONResponse(w, http.StatusInternalServerError, message)
		return
	} else {
		message := images
		utils.SendJSONResponse(w, http.StatusOK, message)
	}
}
