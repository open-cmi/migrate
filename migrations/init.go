package migrations

import (
	"github.com/open-cmi/migrate/cmdopt"
)

// Register register
func Register(seq *cmdopt.SeqInfo) {
	cmdopt.GoMigrationList = append(cmdopt.GoMigrationList, *seq)
}

// Init init
func Init() {

}
