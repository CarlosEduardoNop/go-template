package main

import (
	"flag"
	"os"

	"GoTemplate/pkg/migration"
)

func main() {
	makeTypeCmd := flag.NewFlagSet("type", flag.ExitOnError)

	switch os.Args[1] {
	case "migration":
		migrationName := makeTypeCmd.String("name", "", "Name of migration")
		makeTypeCmd.Parse(os.Args[2:])
		migration.Make(migrationName)
	case "migrate":
		migration.Migrate()
	}
}
