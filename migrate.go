package migrate

import (
	"os"
	"sort"
	"strings"

	"github.com/open-cmi/migrate/cmdopt"
)

// Register register
func Register(seq *cmdopt.SeqInfo) {
	cmdopt.GoMigrationList = append(cmdopt.GoMigrationList, *seq)
	sort.SliceStable(cmdopt.GoMigrationList, func(i, j int) bool {
		cmp := strings.Compare(cmdopt.GoMigrationList[i].Seq, cmdopt.GoMigrationList[j].Seq)
		return cmp == -1
	})
}

// Run run command
func Run(service string) {
	// init service
	cmdopt.Service = service

	// parse command
	opt := cmdopt.ParseArgs()
	opt.Run()
}

func IsSubCommand(cmd string) bool {
	for _, c := range cmdopt.SubCommands {
		if c == cmd {
			return true
		}
	}
	return false
}

// TryRun 尝试运行，如果不是migrate的命令，则返回false
func TryRun(service string) bool {
	if len(os.Args) > 1 && IsSubCommand(os.Args[1]) {
		Run(service)
		return true
	}
	return false
}
