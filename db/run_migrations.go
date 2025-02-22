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

func RunMigrations(config *ui.Config, logger *log.Logger, conn *sql.DB) error {
	ctx := context.Background()

	unapplied_migrations := UnappliedMigrations(config.MigrationsPath, ctx, conn, logger)
	logger.Debug("found the following unapplied migrations:", unapplied_migrations)

	if len(unapplied_migrations) == 0 {
		logger.Print("No migrations to run")
		return nil
	}

	logger.Print("Running migrations...")
	for _, migration := range unapplied_migrations {
		logger.Debug("applying migration ", migration.Id)
		err := migration.Run(ctx, conn, logger)
		if err != nil {
			return err
		}
	}
	logger.Print("Finished migrations")

	return nil
}
