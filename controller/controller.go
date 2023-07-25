package controller

import (
	"gosftp/config"
	"gosftp/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getSftpConfigFromHeaders(c *gin.Context) *config.SftpConfig {
    host := c.GetHeader("sftp-host")
    port := c.GetHeader("sftp-port")
    username := c.GetHeader("sftp-username")
    password := c.GetHeader("sftp-password")

    return &config.SftpConfig{
        Host:     host,
        Port:     port,
        Username: username,
        Password: password,
    }
}

func UploadFile(c *gin.Context) {
    fileHeader, err := c.FormFile("file")
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    file, err := fileHeader.Open()
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    defer file.Close()

    remoteDir := c.PostForm("remoteDir")

    sftpConfig := getSftpConfigFromHeaders(c)
    sftpService, err := service.NewSftpService(sftpConfig)
	if err != nil {
		log.Printf("Failed to create SFTP service: %s", err)
		// Return or handle the error properly
		return
	}

    err = sftpService.UploadFile(file, fileHeader.Filename, remoteDir)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully"})
}