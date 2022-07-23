package main

import (
	"time"

	"github.com/dev-homies/first-project/api/core"
	docs "github.com/dev-homies/first-project/api/docs"
	"github.com/dev-homies/first-project/api/routes/v1"

	"github.com/gin-contrib/cors"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	core.SetupLogger()
	core.SetupConfig()

	core.SetupDatabase()
	core.ProvisionDatabase()

	server := SetupServer()
	server.Run(":4000")
}

func AddRoutes(server *gin.Engine) {
	docs.SwaggerInfo.Host = "http://localhost:4000"
	docs.SwaggerInfo.BasePath = "/"
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	v1 := server.Group("/v1")
	v1.GET("/", routes.Index)
	v1.POST("/register", routes.Register)
}

func SetupServer() *gin.Engine {
	server := gin.Default()
	server.Use(ginzap.Ginzap(core.Logger, time.RFC3339, true))
	server.Use(ginzap.RecoveryWithZap(core.Logger, true))
	server.Use(configureCORS())

	AddRoutes(server)
	return server
}

func configureCORS() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:4000"},
		AllowMethods:     []string{"POST", "GET", "PUT", "DELETE"},
		AllowHeaders:     []string{"Authorization", "Content-Type", "X-Requested-With", "User-Agent"},
		ExposeHeaders:    []string{"Content-Range", "Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc:  func(origin string) bool { return true },
		MaxAge:           12 * time.Hour,
	})
}
