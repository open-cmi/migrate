package main

import (
	"github.com/open-cmi/migrate"
	_ "github.com/open-cmi/migrate/main/migration"
)

func main() {
	migrate.Run("example")
}
