package migrate

import (
	"fmt"

	"github.com/open-cmi/migrate/config"

	"github.com/open-cmi/migrate/cmdopt"
)

// Init module init
func Init() {

}

// SetConfigFile set config file
func SetConfigFile(configfile string) {
	err := config.Init(configfile)
	if err != nil {
		fmt.Printf("init config failed: %s\n", err.Error())
		return
	}
}

// SetMigrateMode set migrate mode
func SetMigrateMode(mode string) {
	cmdopt.SetMigrateMode(mode)
}

// SetMigrateDir set migrate dir
func SetMigrateDir(dir string) {
	cmdopt.SetMigrateDir(dir)
}

// Register register
func Register(seq *cmdopt.SeqInfo) {
	cmdopt.GoMigrationList = append(cmdopt.GoMigrationList, *seq)
}

// Run run command
func Run() {
	opt := cmdopt.ParseArgs()
	opt.Run()
}
