package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"spotify/config"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func dsn(cfg *config.DBConfig) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.SSLMode)
}

func New(ctx context.Context, cfg *config.DBConfig) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn(cfg))
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetConnMaxLifetime(cfg.ConnMaxLifetime)

	if err := db.PingContext(ctx); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
