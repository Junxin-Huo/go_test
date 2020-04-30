package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type formA struct {
	Foo string `json:"foo" xml:"foo" binding:"required"`
}

type formB struct {
	Bar string `json:"bar" xml:"bar" binding:"required"`
}

func SomeHandler(c *gin.Context) {
	//objA := formA{}
	//objB := formB{}
	//// This c.ShouldBind consumes c.Request.Body and it cannot be reused.
	//if errA := c.ShouldBind(&objA); errA == nil {
	//	c.String(http.StatusOK, `the body should be formA`)
	//	// Always an error is occurred by this because c.Request.Body is EOF now.
	//} else if errB := c.ShouldBind(&objB); errB == nil {
	//	c.String(http.StatusOK, `the body should be formB`)
	//} else {
	//	c.String(http.StatusOK, `???`)
	//}

	//objA := formA{}
	//if errA := c.ShouldBind(&objA); errA == nil {
	//	c.String(http.StatusOK, `the body should be formA`)
	//	// Always an error is occurred by this because c.Request.Body is EOF now.
	//}

	objB := formB{}
	if errB := c.ShouldBind(&objB); errB == nil {
		c.String(http.StatusOK, `the body should be formB`)
	}
}

func main() {
	router := gin.Default()

	router.GET("/ping", SomeHandler)

	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
