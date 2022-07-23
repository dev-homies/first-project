package routes

import (
	"net/http"

	"github.com/dev-homies/first-project/api/core"
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
	core.Logger.Debug("Recieved ping on '/'")
	response := IndexResponse{Body: "Hello world!"}

	core.RequestsProcessedMetric.Inc()
	c.JSON(http.StatusOK, response)
}
