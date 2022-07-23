package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type IndexResponse struct {
	Body string `json:"body"`
}

// @Summary index
// @Description Index
// @Accept json
// @Produce json
// @Success 200 {string} Index
// @Router /v1/ [get]
func Index(c *gin.Context) {
	response := IndexResponse{Body: "Hello world!"}
	c.JSON(http.StatusOK, response)
}
