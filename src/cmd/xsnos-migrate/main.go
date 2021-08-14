package main

import (
	"flag"
	"fmt"
	"path/filepath"

	"github.com/open-cmi/goutils/common"
	"github.com/open-cmi/goutils/database"
	"github.com/open-cmi/goutils/database/dbsql"
	"github.com/open-cmi/migrate/cmdopt"
	"github.com/open-cmi/migrate/config"
	"github.com/open-cmi/migrate/global"
	"github.com/open-cmi/migrate/migrations"
)

var configfile string = ""

func main() {

	flag.StringVar(&configfile, "config", configfile, "config file")
	flag.Parse()

	if configfile == "" {
		rp := common.GetRootPath()
		configfile = filepath.Join(rp, "xsnos-migrate", "etc", "db.json")
	}
	err := config.Init(configfile)
	if err != nil {
		fmt.Printf("init config failed: %s\n", err.Error())
		return
	}

	conf := config.GetConfig()
	var dbconf database.Config
	dbconf.Type = conf.Model.Type
	if dbconf.Type == "sqlite3" {
		dbconf.File = conf.Model.File
	} else {
		dbconf.Host = conf.Model.Address
		dbconf.Port = conf.Model.Port
		dbconf.User = conf.Model.User
		dbconf.Password = conf.Model.Password
		dbconf.Database = conf.Model.Database
	}

	db, err := dbsql.SQLInit(&dbconf)
	if err != nil {
		fmt.Printf("db init failed: %s\n", err.Error())
		return
	}
	global.DB = db

	migrations.Init()
	opt := cmdopt.ParseArgs()
	opt.Run()
}
