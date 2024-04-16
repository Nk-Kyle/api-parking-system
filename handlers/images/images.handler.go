package images

import (
	"net/http"
	"os"

	"api-parking-system/gcs"

	"github.com/gin-gonic/gin"
)

// Image Upload godoc
// @Summary Upload an image
// @Description Upload an image
// @Tags images
// @Accept  json
// @Produce  json
// @Param image formData file true "Image file"
// @Success 201 {string} string "Image Url"
func Upload(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Image is required",
		})
		return
	}

	uploader := gcs.FileUploader{
		Client:     gcs.StorageClient,
		ProjectID:  os.Getenv("GCS_PROJECT_ID"),
		BucketName: os.Getenv("GCS_BUCKET_NAME"),
		UploadPath: os.Getenv("GCS_UPLOAD_PATH"),
	}

	blob, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error opening file",
		})
		return
	}

	url, err := uploader.UploadFile(blob, file.Filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error uploading image",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"url": url,
	})
}
