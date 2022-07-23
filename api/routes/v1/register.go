package routes

import (
	"context"
	"net/http"

	"github.com/dev-homies/first-project/api/core"
	"github.com/dev-homies/first-project/api/models"
	"github.com/gin-gonic/gin"
)

// @Summary register
// @Description Register
// @Accept json
// @Produce json
// @Success 200 {string} Index
// @Router /v1/register [post]
func Register(c *gin.Context) {
	user := models.User{}

	if err := c.ShouldBindJSON(&user); err != nil {
		core.Logger.Sugar().Warnf("Should bind to JSON error: %v", err)
		c.JSON(http.StatusUnauthorized, "Invalid json provided")
		return
	}

	userInfo := &models.User{
		Name:     user.Name,
		Password: user.Password,
	}

	res, err := core.Database.NewInsert().Model(userInfo).Exec(context.Background())
	if err != nil {
		core.Logger.Sugar().Warnf("Insert error: %v", err)
		c.JSON(http.StatusUnauthorized, "Cannot input data.")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"userInfo": res,
	})
}