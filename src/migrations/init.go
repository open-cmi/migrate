package migrations

import (
	"github.com/open-cmi/migrate/cmdopt"
)

// Register register
func Register(seq *cmdopt.SeqInfo) {
	cmdopt.MigrationList = append(cmdopt.MigrationList, *seq)
}

// Init init
func Init() {

}
