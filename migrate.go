package migrate

import (
	"github.com/open-cmi/migrate/cmdopt"
)

// Init module init
func Init(service string) error {

	cmdopt.Service = service
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

func IsMigrateCommand(cmd string) bool {
	for _, c := range cmdopt.SubCommands {
		if c == cmd {
			return true
		}
	}
	return false
}
