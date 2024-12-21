package db

import (
	"my-project/src/config"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var dbPool *pgxpool.Pool

// InitDb initializes the database connection using pgxpool with custom settings
func InitDb(cfg *config.Config) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cnn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s&timezone=Asia/Tehran",
		cfg.Postgres.User, cfg.Postgres.Password, cfg.Postgres.Host, cfg.Postgres.Port,
		cfg.Postgres.DbName, cfg.Postgres.SSLMode)

	config, err := pgxpool.ParseConfig(cnn)
	if err != nil {
		return fmt.Errorf("failed to parse database config: %w", err)
	}

	// Set custom connection settings
	config.MaxConns = int32(cfg.Postgres.MaxOpenConns)                 // Max open connections
	config.MinConns = int32(cfg.Postgres.MaxIdleConns)                 // Min idle connections
	config.MaxConnLifetime = time.Duration(cfg.Postgres.ConnMaxLifetime) * time.Minute // Connection max lifetime
	config.MaxConnIdleTime = 5 * time.Minute                           // Optional: idle timeout
	config.HealthCheckPeriod = 1 * time.Minute                         // Optional: health check interval

	// Create connection pool
	dbPool, err = pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return fmt.Errorf("failed to create pgxpool: %w", err)
	}

	// Check connection
	err = dbPool.Ping(ctx)
	if err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("Db connection established with custom settings")
	return nil
}

// GetDb returns the database pool
func GetDb() *pgxpool.Pool {
	return dbPool
}

// CloseDb closes the database connection pool
func CloseDb() {
	if dbPool != nil {
		dbPool.Close()
		log.Println("Db connection closed")
	}
}
