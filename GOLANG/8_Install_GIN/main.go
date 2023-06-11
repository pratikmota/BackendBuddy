package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HelloFunc(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello backend buddy",
	})
}

// go get -u github.com/gin-gonic/gin
func main() {
	instance := gin.Default()
	instance.GET("/hello", HelloFunc)
	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	instance.Run()
}
