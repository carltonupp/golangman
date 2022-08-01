package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/api/word", GetWordHandler)

	r.Run()
}

func GetWordHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"word": "example",
	})
}
