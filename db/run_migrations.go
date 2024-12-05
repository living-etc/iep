package db

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/charmbracelet/log"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
	_ "modernc.org/sqlite"

	"github.com/living-etc/iep/ui"
)

func InitDb(ctx context.Context, filename string) (*sql.DB, error) {
	conn, err := sql.Open("libsql", "file:"+filename)
	if err != nil {
		return nil, err
	}

	err = conn.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	Exec(ctx, conn, "CREATE TABLE IF NOT EXISTS migrations(id TEXT NOT NULL, PRIMARY KEY(id))")

	return conn, nil
}

func Exec(ctx context.Context, conn *sql.DB, statement string, args ...any) sql.Result {
	res, err := conn.ExecContext(ctx, statement, args...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to execute statement %s: %s", statement, err)
		os.Exit(1)
	}

	return res
}

func RunMigrations(config ui.Config, logger *log.Logger) error {
	ctx := context.Background()

	conn, err := InitDb(ctx, config.ExerciseDatabase)
	if err != nil {
		return err
	}
	defer conn.Close()

	unapplied_migrations := Get(ctx, conn)

	if len(unapplied_migrations) == 0 {
		logger.Print("No migrations to run")
		return nil
	}

	logger.Print("Running migrations...")
	for _, migration := range unapplied_migrations {
		err := migration.Run(ctx, conn)
		if err != nil {
			return err
		}
	}
	logger.Print("Finished migrations")

	return nil
}
