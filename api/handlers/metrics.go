package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// @Summary metrics
// @Description Prometheus metrics
// @Produce html
// @Success 200 {string} Metrics
// @Router /metrics [get]
func MetricsHandler() gin.HandlerFunc {
	handler := promhttp.Handler()

	return func(c *gin.Context) {
		handler.ServeHTTP(c.Writer, c.Request)
	}
}
