package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type IndexResponse struct {
	Body string `json:"body"`
}

func index(c *gin.Context) {
	response := IndexResponse{Body: "Hello world!"}
	c.JSON(http.StatusOK, response)
}

func main() {
	r := gin.Default()
	r.GET("/", index)
	r.Run(":4000")
}
