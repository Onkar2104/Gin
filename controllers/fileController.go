package controllers

import (
	"fmt"
	"net/http"
	"time"

	"path/filepath"

	"github.com/Onkar2104/go/initializers"
	"github.com/Onkar2104/go/models"
	"github.com/gin-gonic/gin"
)

// UploadFile handles file upload requests
func UploadFile(c *gin.Context) {
	// Log request to check if the file field exists
	fmt.Println("Receiving upload request...")

	// Parse the uploaded file
	file, err := c.FormFile("file")
	if err != nil {
		fmt.Println("Error retrieving file:", err) // Debugging log
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file found in request"})
		return
	}

	// Define upload path
	uploadPath := "uploads/" + filepath.Base(file.Filename)

	// Save file
	if err := c.SaveUploadedFile(file, uploadPath); err != nil {
		fmt.Println("Error saving file:", err) // Debugging log
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	fileRecord := models.File{
		Name:      file.Filename,
		Path:      uploadPath,
		CreatedAt: time.Now(), // Set timestamp
	}
	result := initializers.DB.Create(&fileRecord)
	if result.Error != nil {
		fmt.Println("Error inserting into database:", result.Error) // Debugging log
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store file in database"})
		return
	}

	// Success response
	c.JSON(http.StatusOK, gin.H{
		"message": "File uploaded successfully!",
		"file":    uploadPath,
	})

	// c.Redirect(http.StatusFound, "/view/upload")
}

// GetFiles fetches all uploaded files
func GetFiles(c *gin.Context) {
	var files []models.File
	initializers.DB.Find(&files)
	c.JSON(http.StatusOK, files)
}

func ShowUploadPage(c *gin.Context) {
	c.HTML(http.StatusOK, "upload.html", nil)
}


