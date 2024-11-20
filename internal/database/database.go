package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func InitDB() *pgxpool.Pool {
	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("DB_USER"), os.Getenv("DB_PWD"), os.Getenv("DB_DOMAIN"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))

	config, err := pgxpool.ParseConfig(dbUrl)

	if err != nil {
		panic("Error parsing db config: " + err.Error())
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		panic("Error creating db pool: " + err.Error())
	}

	return pool
}
