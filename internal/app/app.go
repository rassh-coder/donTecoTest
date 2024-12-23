package app

import (
	"context"
	"donTecoTest/config"
	"donTecoTest/pkg/handler"
	"donTecoTest/pkg/httpserver"
	"donTecoTest/pkg/repository"
	"donTecoTest/pkg/service"
	"fmt"
	"github.com/jackc/pgx/v5"
)

func Run(cfg *config.Config) error {
	db, err := repository.NewPostgresDB(cfg)
	if err != nil {
		return err
	}

	defer func(db *pgx.Conn, ctx context.Context) {
		err := db.Close(ctx)
		if err != nil {
			panic(fmt.Sprintf("can't close db connection: %s", err))
		}
	}(db, context.Background())

	r := repository.NewRepository(db)
	s := service.NewService(r)
	h := handler.NewHandler(s)

	err = httpserver.Server(h.InitRoutes(), cfg)
	if err != nil {
		return err
	}

	return nil
}
