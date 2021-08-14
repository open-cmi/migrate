package config

import (
	"errors"
	"fmt"

	"github.com/open-cmi/goutils/confparser"
	"github.com/open-cmi/goutils/database"
	"github.com/open-cmi/goutils/database/dbsql"
	"github.com/open-cmi/migrate/global"
)

// ConfParser conf parser
var ConfParser *confparser.Parser

// DatabaseModel database model
type DatabaseModel struct {
	Type     string `json:"type"`
	File     string `json:"file,omitempty"`
	Address  string `json:"address,omitempty"`
	Port     int    `json:"port,omitempty"`
	Database string `json:"database,omitempty"`
	User     string `json:"user,omitempty"`
	Password string `json:"password,omitempty"`
}

// Config config structure
type Config struct {
	Model DatabaseModel `json:"model"`
}

var config Config

// InitDB init db
func InitDB() error {
	var dbconf database.Config
	dbconf.Type = config.Model.Type
	if dbconf.Type == "sqlite3" {
		dbconf.File = config.Model.File
	} else {
		dbconf.Host = config.Model.Address
		dbconf.Port = config.Model.Port
		dbconf.User = config.Model.User
		dbconf.Password = config.Model.Password
		dbconf.Database = config.Model.Database
	}

	db, err := dbsql.SQLInit(&dbconf)
	if err != nil {
		fmt.Printf("db init failed: %s\n", err.Error())
		return err
	}
	global.DB = db
	return nil
}

// Init config module init
func Init(configfile string) (err error) {
	parser := confparser.New(configfile)
	if parser == nil {
		return errors.New("parse config failed")
	}
	err = parser.Load(&config)
	ConfParser = parser
	return InitDB()
}
