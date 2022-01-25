package main

import (
	"github.com/open-cmi/migrate"
	
)

func main() {

	migrate.Init("example")
	// if you use go mode, should init here

	migrate.Run()
}
