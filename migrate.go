package migrate

import (
	"github.com/open-cmi/migrate/cmdopt"
)

// Init module init
func Init() error {

	return nil
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
