package main

import (
	"github.com/open-cmi/migrate"
)

var configfile string = ""
var migratedir string = ""

func main() {

	migrate.Init()
	// if you use go mode, should init here

	migrate.Run()
}
