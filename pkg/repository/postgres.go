package repository

import (
	"context"
	"donTecoTest/config"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
)

var dbLink *pgx.Conn

func NewPostgresDB(cfg *config.Config) (*pgx.Conn, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)
	conn, err := pgx.Connect(context.Background(), connStr)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("can't connect to database: %s", err))
	}

	dbLink = conn

	return conn, nil
}

func GetDB() (*pgx.Conn, error) {
	if dbLink == nil {
		return nil, errors.New("no db connection")
	}

	return dbLink, nil
}
