package schema

import (
	"embed"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

func Migrate(fs *embed.FS, psqlURL string) {
	source, err := iofs.New(fs, "migrations")
	if err != nil {
		log.Fatalf("Failed to read migrations source: %s", err)

		return
	}

	instance, err := migrate.NewWithSourceInstance("iofs", source, psqlURL)
	if err != nil {
		log.Fatalf("Failed to initialization the migrate instance: %s", err)

		return
	}

	err = instance.Up()

	switch err {
	case nil:
		log.Println("The migration schema: The schema successfully upgraded!")
	case migrate.ErrNoChange:
		log.Println("The migration schema: The schema not changed")
	default:
		log.Fatalf("Could not apply the migration schema: %s", err)
	}
}
