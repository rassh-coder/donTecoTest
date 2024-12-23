package migrationDown

import (
	"context"
	"github.com/jackc/pgx/v5"
	"log"
)

func Run(db *pgx.Conn) error {
	_, err := db.Exec(context.Background(), `
		DROP INDEX IF EXISTS name_idx;
	`)
	if err != nil {
		return err
	}

	_, err = db.Exec(context.Background(), `
		DROP TABLE IF EXISTS employees;
	`)

	if err != nil {
		return err
	}
	defer func(db *pgx.Conn, ctx context.Context) {
		err := db.Close(ctx)
		if err != nil {
			log.Fatalf("can't close db connection: %s", err)
		}
	}(db, context.Background())
	return nil
}
