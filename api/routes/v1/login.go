package routes

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/dev-homies/first-project/api/core"
	"github.com/dev-homies/first-project/api/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	Name string `json:"name"`
	jwt.RegisteredClaims
}

func Login(c *gin.Context) {
	user := models.User{}

	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Printf("Should bind to JSON error: %v", err)

		c.JSON(http.StatusUnauthorized, "Invalid json provided")
		return
	}

	userInfo := &models.User{
		Name:     user.Name,
	}

	user1 := new(models.User)
	err := core.Database.NewSelect().Model(user1).Where("Name = ?", userInfo.Name).Scan(context.Background())
	if err != nil {
		fmt.Printf("Insert error: %v", err)

		c.JSON(http.StatusUnauthorized, "Cannot find user.")
		return
	}

	expirationTime := time.Now().Add(15 * time.Minute)
	claims := &Claims{
		Name: userInfo.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			// set the expire time
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("my_secret_key"))
	if err != nil {
		fmt.Printf("Insert error: %v", err)

		c.JSON(http.StatusUnauthorized, "Cannot create access token.")
		return
	}

	c.SetCookie("accessToken", tokenString, 3600000, "/login", "localhost", false, false)
}