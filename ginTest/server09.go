package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	// Simple group: v1
	v1 := router.Group("/v1")
	{
		v1.POST("/login", func(c *gin.Context) {
			c.String(http.StatusOK, "v1 loginEndpoint")
		})
		v1.POST("/submit", func(c *gin.Context) {
			c.String(http.StatusOK, "v1 submitEndpoint")
		})
		v1.POST("/read", func(c *gin.Context) {
			c.String(http.StatusOK, "v1 readEndpoint")
		})
	}

	// Simple group: v2
	v2 := router.Group("/v2")
	{
		v2.POST("/login", func(c *gin.Context) {
			c.String(http.StatusOK, "v2 loginEndpoint")
		})
		v2.POST("/submit", func(c *gin.Context) {
			c.String(http.StatusOK, "v2 submitEndpoint")
		})
		v2.POST("/read", func(c *gin.Context) {
			c.String(http.StatusOK, "v2 readEndpoint")
		})
	}

	router.POST("/v1/loginn", func(c *gin.Context) {
		c.String(http.StatusOK, "aaa")
	})

	router.POST("/v1/v111", func(c *gin.Context) {
		c.String(http.StatusOK, "vvvv")
	})

	router.Run(":8080")
}
