package db

import (
	"context"
	"database/sql"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/charmbracelet/log"
)

func UnappliedMigrations(
	migrationsDirectory string,
	ctx context.Context,
	db *sql.DB,
	logger *log.Logger,
) []Migration {
	unapplied_migrations := []Migration{}

	migration_files, err := os.ReadDir(migrationsDirectory)
	if err != nil {
		panic(err)
	}

	migration_ids := []string{}
	for _, migration_file := range migration_files {
		filename := filepath.Base(migration_file.Name())
		migration_id := strings.TrimSuffix(filename, filepath.Ext(filename))
		migration_ids = append(migration_ids, migration_id)
	}
	sort.Strings(migration_ids)

	for _, migration_id := range migration_ids {
		var migration string
		if err := db.QueryRow("SELECT id FROM migrations WHERE id = ?", migration_id).Scan(&migration); err != nil {
			if err == sql.ErrNoRows {
				unapplied_migrations = append(
					unapplied_migrations,
					Migration{
						Id:       migration_id,
						Filepath: filepath.Join(migrationsDirectory, migration_id+".sql"),
					},
				)
			}
		}
	}

	return unapplied_migrations
}
