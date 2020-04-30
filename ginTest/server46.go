package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/v1/topics", func(c *gin.Context) {
		if c.Query("username") == "" {
			c.String(http.StatusOK, "topics")
		} else {
			c.String(http.StatusOK, "username=%s", c.Query("username"))
		}
	})

	router.GET("/v1/topics/:topic_id", func(c *gin.Context) {
		c.String(http.StatusOK, "topic=%s", c.Param("topic_id"))
	})

	router.Run()
}
