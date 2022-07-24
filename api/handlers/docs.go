package handlers

import (
	"github.com/dev-homies/first-project/api/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @Summary docs
// @Description Swagger docs
// @Produce html
// @Success 200 {string} Docs
// @Router /swagger/index.html [get]
func DocsHandler() gin.HandlerFunc {
	docs.SwaggerInfo.Host = "http://localhost:4000"
	docs.SwaggerInfo.BasePath = "/"

	return ginSwagger.WrapHandler(swaggerfiles.Handler)
}
