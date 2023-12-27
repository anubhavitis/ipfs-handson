package handlers

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func init() {
	os.MkdirAll(storageDir, os.ModePerm)
}

func UploadFileHandler(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// To generate random UUID for the file.
	fileId := uuid.New().String() + filepath.Ext(file.Filename)
	filePath := filepath.Join(storageDir, fileId)

	// Save the file in local
	// Todo: Add push to IPFS
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"fileId": fileId,
	})
}

func GetFileHandler(c *gin.Context) {
	fileId := c.Param("fileId")
	filePath := filepath.Join(storageDir, fileId)

	// Check if the file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}

	// Serve the file
	c.File(filePath)
}
