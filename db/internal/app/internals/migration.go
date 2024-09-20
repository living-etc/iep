package internals

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Migration struct {
	Filepath string
}

func (m *Migration) Run(ctx context.Context, db *sql.DB) error {
	statement, err := os.ReadFile(m.Filepath)
	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stdout, "Running migrations...\n")
	migration_id := strings.TrimSuffix(filepath.Base(m.Filepath), filepath.Ext(m.Filepath))
	fmt.Fprintf(os.Stdout, "\t%s\n", migration_id)
	fmt.Fprintf(os.Stdout, "Finished migrations\n")

	exec(ctx, db, string(statement))

	add_migration_id_statement := fmt.Sprintf(addMigrationIdSQL, migration_id)
	exec(ctx, db, add_migration_id_statement)

	return nil
}
