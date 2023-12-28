package handlers

import (
	IPFS "fileverse-test/src/services/ipfs"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func UploadFileHandler(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	uploadedFile, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer uploadedFile.Close()
	cid, _ := IPFS.AddFile(file.Filename, &uploadedFile)

	c.JSON(http.StatusOK, gin.H{
		"fileId": cid,
	})
}

func GetFileHandler(c *gin.Context) {
	fileId := c.Param("fileId")
	filePath, err := IPFS.GetFile(fileId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.File(filePath)
	os.Remove(filePath)
}
