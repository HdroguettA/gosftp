package main

import (
	"gosftp/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	sftpGroup := r.Group("/sftp")
	{
		sftpGroup.POST("/upload", controller.UploadFile)
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not Found"})
	})

	r.Run()
}
