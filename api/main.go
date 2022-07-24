package main

import (
	"time"

	"github.com/dev-homies/first-project/api/core"
	"github.com/dev-homies/first-project/api/handlers"
	"github.com/dev-homies/first-project/api/middleware"
	"github.com/dev-homies/first-project/api/routes/v1"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
)

// NOTE: We might be better storing "globals" in the Gin context instead.
func main() {
	core.SetupConfig()
	core.SetupLogger()

	core.SetupDatabase()
	core.ProvisionDatabase()

	server := SetupServer()
	server.Run(":4000")
}

func AddRoutes(r *gin.Engine) {
	r.GET("/swagger/*any", handlers.DocsHandler())
	r.GET("/metrics", handlers.MetricsHandler())

	v1 := r.Group("/v1")
	v1.GET("/", routes.Index)
	v1.POST("/register", routes.Register)
}

func SetupServer() *gin.Engine {
	r := gin.Default()
	r.Use(ginzap.Ginzap(core.RawLogger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(core.RawLogger, true))
	r.Use(middleware.CorsMiddleware())

	AddRoutes(r)
	return r
}
