package main

import (
	"errors"
	"flag"
	"fmt"

	// Библиотека для миграций
	"github.com/golang-migrate/migrate/v4"
	// Драйвер для выполнения миграций с SQLite3 без CGO
	_ "github.com/golang-migrate/migrate/v4/database/sqlite"
	// Дравйвер для получения миграций из файлов
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	var storagePath, migrationsPath, migrationTable string

	flag.StringVar(&storagePath, "storage-path", "", "Path to a directory containing the migration files")
	flag.StringVar(&migrationsPath, "migrations-path", "", "Path to a directory containing the migration files")
	flag.StringVar(&migrationTable, "migration-table", "migrations", "Name of the migration table")
	flag.Parse()

	if storagePath == "" {
		panic("storage-path is required")
	}

	if migrationsPath == "" {
		panic("migrations-path is required")
	}

	m, err := migrate.New(
		"file://"+migrationsPath,
		fmt.Sprintf("sqlite://%s?x-migrations-table=%s", storagePath, migrationTable),
	)
	if err != nil {
		panic(err)
	}

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("No migrations to apply")

			return
		}

		panic(err)
	}

	fmt.Println("Applied migrations successfully")
}
