package main

import (
	"fmt"

	"github.com/open-cmi/goutils/config"
	"github.com/open-cmi/goutils/database"
	"github.com/open-cmi/goutils/database/dbsql"
	"github.com/open-cmi/migrate/cmdopt"
	"github.com/open-cmi/migrate/global"
)

func main() {
	conf, err := config.InitConfig()
	if err != nil {
		fmt.Printf("init config failed\n")
		return
	}
	var dbconf database.Config
	dbconf.Type = conf.GetStringMap("model")["type"].(string)
	if dbconf.Type == "sqlite3" {
		dbconf.File = conf.GetStringMap("model")["location"].(string)
	} else {
		dbconf.Host = conf.GetStringMap("model")["host"].(string)
		dbconf.Port = conf.GetStringMap("model")["port"].(int)
		dbconf.User = conf.GetStringMap("model")["user"].(string)
		dbconf.Password = conf.GetStringMap("model")["password"].(string)
		dbconf.Database = conf.GetStringMap("model")["database"].(string)
	}

	db, err := dbsql.SQLInit(&dbconf)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}
	global.Conf = conf
	global.DB = db

	opt := cmdopt.ParseArgs()
	opt.Run()
}
