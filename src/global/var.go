package global

import (
	"database/sql"

	"github.com/spf13/viper"
)

var Conf *viper.Viper
var DB *sql.DB
