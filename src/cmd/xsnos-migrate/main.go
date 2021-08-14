package main

import (
	"flag"
	"path/filepath"

	"github.com/open-cmi/goutils/common"
	"github.com/open-cmi/migrate"
)

var configfile string = ""
var migratedir string = ""

func main() {
	flag.StringVar(&configfile, "config", configfile, "config file")
	flag.StringVar(&migratedir, "migrate-dir", migratedir, "migration directory")
	flag.Parse()

	if configfile == "" {
		rp := common.Getwd()
		configfile = filepath.Join(rp, "etc", "db.json")
	}

	migrate.Init()
	migrate.SetConfigFile(configfile)
	if migratedir == "" {
		migrate.SetMigrateMode("go")
		// 如果是go脚本，需要在这里初始化
	} else {
		migrate.SetMigrateMode("sql")
		if migratedir != "" {
			migrate.SetMigrateDir(migratedir)
		}
	}

	// run command
	migrate.Run()
}
