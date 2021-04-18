package main

import (
	"fmt"

	"github.com/open-cmi/goutils/database"
	"github.com/open-cmi/goutils/database/dbsql"
	"github.com/open-cmi/migrate/cmdopt"
	"github.com/open-cmi/migrate/config"
	"github.com/open-cmi/migrate/global"
)

func main() {
	err := config.Init()
	if err != nil {
		fmt.Printf("init config failed\n")
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
		dbconf.Database = conf.Model.DB
	}

	db, err := dbsql.SQLInit(&dbconf)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}
	global.DB = db

	opt := cmdopt.ParseArgs()
	opt.Run()
}
