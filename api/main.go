package main

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type IndexResponse struct {
	Body string `json:"body"`
}

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	ID       int64 `bun:",pk,autoincrement,type:uuid"`
	Name     string
	Password string
}

func index(c *gin.Context) {
	response := IndexResponse{Body: "Hello world!"}
	c.JSON(http.StatusOK, response)
}

func main() {
	CreateTables(context.Background())

	// Create gin server
	r := gin.Default()
	// Add version to api, set to a group
	v1 := r.Group("/v1")
	v1.GET("/", index)

	r.Run(":4000")
}

func CreateTables(ctx context.Context) {

	dsn := "postgres://postgres:dev@localhost:5432/firstproject?sslmode=disable"
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	db := bun.NewDB(sqldb, pgdialect.New())
	_, err := db.NewCreateTable().IfNotExists().
		Model((*User)(nil)).Exec(ctx)
	if err != nil {
		panic(err)
	}
}
