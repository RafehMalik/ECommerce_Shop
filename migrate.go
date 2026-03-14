package main

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	dbURL := "postgres://postgres:password@localhost:5432/ecommerce_shop?sslmode=disable"

	m, err := migrate.New("file://db/migrations", dbURL)
	if err != nil {
		log.Fatal(err)
	}

	if err := m.Up(); err != nil {
		if err.Error() == "no change" {
			log.Println("All migrations already applied")
		} else {
			log.Fatal(err)
		}
	} else {
		log.Println("Migrations applied successfully!")
	}
}
