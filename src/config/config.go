package config

import (
	"errors"
	"path/filepath"

	"github.com/open-cmi/goutils/common"
	"github.com/open-cmi/goutils/confparser"
)

var ConfParser *confparser.Parser

type DatabaseModel struct {
	Type     string `json:"type"`
	File     string `json:"file"`
	Address  string `json:"address"`
	Port     int    `json:"port"`
	DB       string `json:"db"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type Config struct {
	Model DatabaseModel `json:"model"`
}

var config Config

func Init() (err error) {
	rp := common.GetRootPath()
	configfile := filepath.Join(rp, "etc", "config.json")

	parser := confparser.New(configfile)
	if parser == nil {
		return errors.New("parse config failed")
	}
	err = parser.Load(&config)
	ConfParser = parser
	return err
}

func Save(c *Config) {
	ConfParser.Save(c)
}

func GetConfig() *Config {
	return &config
}
