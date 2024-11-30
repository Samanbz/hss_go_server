package test

import (
	"context"
	"fmt"
	"hss/internal/api/routes"
	"hss/pkg/utils/singleton"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

func SetupTestContainer(ctx context.Context) (*postgres.PostgresContainer, *pgxpool.Pool, *fiber.App, error) {

	container, err := postgres.Run(ctx,
		"postgres:17",
		postgres.WithDatabase("testdb"),
		postgres.WithUsername("testuser"),
		postgres.WithPassword("testpassword"),
		postgres.WithInitScripts("sql/init_schema.sql"),
		testcontainers.WithWaitStrategy(wait.ForListeningPort("5432/tcp").WithStartupTimeout(30*time.Second)),
	)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to start postgres container: %w", err)
	}

	connStr, err := container.ConnectionString(ctx)
	if err != nil {
		container.Terminate(ctx)
		return nil, nil, nil, fmt.Errorf("failed to get connection string: %w", err)
	}

	config, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		container.Terminate(ctx)
		return nil, nil, nil, fmt.Errorf("failed to parse config: %w", err)
	}

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		container.Terminate(ctx)
		return nil, nil, nil, fmt.Errorf("failed to create pool: %w", err)
	}

	app := fiber.New()

	RequestHandlers, err := singleton.InitSingletons(pool)
	if err != nil {
		container.Terminate(ctx)
		return nil, nil, nil, fmt.Errorf("failed to init singletons: %w", err)
	}

	routes.InitRoutes(app, RequestHandlers)

	return container, pool, app, nil
}

func TeardownTestContainer(ctx context.Context, container *postgres.PostgresContainer, pool *pgxpool.Pool, app *fiber.App) {
	if app != nil {
		if err := app.Shutdown(); err != nil {
			log.Printf("Failed to shut down app: %v", err)
		}
	}

	if pool != nil {
		pool.Close()
	}

	if container != nil {
		if err := container.Terminate(ctx); err != nil {
			log.Printf("Failed to terminate container: %v", err)
		}
	}
}
