package main

import (
	"errors"
	"flag"
	"fmt"

	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/database/postgres"
)

func main() {
	var migrationsPath, migrationsTable, user, password, host, port string

	flag.StringVar(&migrationsPath, "migrations-path", "", "migrations path")
	flag.StringVar(&migrationsTable, "migrations-table", "", "migrations table")
	flag.StringVar(&user, "user", "", "user")
	flag.StringVar(&password, "password", "", "password")
	flag.StringVar(&host, "host", "localhost", "host")
	flag.StringVar(&port, "port", "5432", "port")

	flag.Parse()

	if migrationsPath == "" {
		panic("migrations path is empty")
	}

	dbUrl := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/b?query?x-migrations-table=%s",
		user,
		password,
		host,
		port,
		migrationsTable,
	)
	m, err := migrate.New(
		"file://"+migrationsPath,
		dbUrl,
	)
	if err != nil {
		panic(err)
	}

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("no migrations to apply")

			return
		}

		panic(err)
	}

	fmt.Println("migrations applied successfully")
}
