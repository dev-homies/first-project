package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type IndexResponse struct {
	Body string `json:"body"`
}

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	ID       int64 `bun:",pk,autoincrement"`
	Name     string
	Password string
}

func index(c *gin.Context) {
	response := IndexResponse{Body: "Hello world!"}
	c.JSON(http.StatusOK, response)
}

func GetDBConnection() *bun.DB {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_DB"),
	)

	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	return bun.NewDB(sqldb, pgdialect.New())
}

func Register(c *gin.Context) {
	db := GetDBConnection()
	user := User{}

	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Printf("Should bind to JSON error: %v", err)

		c.JSON(http.StatusUnauthorized, "Invalid json provided")
		return
	}

	userInfo := &User{
		Name:     user.Name,
		Password: user.Password,
	}

	res, err := db.NewInsert().Model(userInfo).Exec(context.Background())
	if err != nil {
		fmt.Printf("Insert error: %v", err)

		c.JSON(http.StatusUnauthorized, "Cannot input data.")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"userInfo": res,
	})
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	CreateTables(context.Background())

	// Create gin server
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:4000"},
		AllowMethods:     []string{"POST", "GET", "PUT", "DELETE"},
		AllowHeaders:     []string{"authorization", "Content-Type", "X-Requested-With", "User-Agent"},
		ExposeHeaders:    []string{"Content-Range", "Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc:  func(origin string) bool { return true },
		MaxAge:           12 * time.Hour,
	}))

	// Add version to api, set to a group
	v1 := r.Group("/v1")
	v1.GET("/", index)
	v1.POST("/register", Register)

	r.Run(":4000")
}

func CreateTables(ctx context.Context) {
	db := GetDBConnection()

	_, err := db.NewCreateTable().IfNotExists().Model((*User)(nil)).Exec(ctx)
	if err != nil {
		panic(err)
	}
}
