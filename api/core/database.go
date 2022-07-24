package core

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/dev-homies/first-project/api/models"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

var Database *bun.DB

func SetupDatabase() {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		Config.GetString("DATABASE_USER"),
		Config.GetString("DATABASE_PASSWORD"),
		Config.GetString("DATABASE_HOST"),
		Config.GetInt("DATABASE_PORT"),
		Config.GetString("DATABASE_DB"),
	)

	db := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	Database = bun.NewDB(db, pgdialect.New())
}

func ProvisionDatabase() {
	_, err := Database.NewCreateTable().IfNotExists().Model((*models.User)(nil)).Exec(context.Background())
	if err != nil {
		Logger.Fatalw("%v", err)
	}
}
