package main

import (
	"github.com/open-cmi/goutils/config"
	"github.com/open-cmi/migrate/cmdopt"
)

func main() {
	config.InitConfig()

	opt := cmdopt.ParseArgs()
	opt.Run()
}
