package main

import (
	"github.com/open-cmi/migrate"
)

func main() {

	migrate.Init()
	// if you use go mode, should init here

	migrate.Run()
}