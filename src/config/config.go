package config

import (
	"errors"
	"path/filepath"

	"github.com/open-cmi/goutils/common"
	"github.com/open-cmi/goutils/confparser"
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

var configfile string = ""

// Init config module init
func Init() (err error) {
	parser := confparser.New(configfile)
	if parser == nil {
		return errors.New("parse config failed")
	}
	err = parser.Load(&config)
	ConfParser = parser
	return err
}

// SetConfigFile set config file before Init
func SetConfigFile(file string) {
	configfile = file
}

// Save save config
func Save(c *Config) {
	ConfParser.Save(c)
}

// GetConfig get config
func GetConfig() *Config {
	return &config
}

func init() {
	rp := common.GetRootPath()
	configfile = filepath.Join(rp, "etc", "db.json")
}
